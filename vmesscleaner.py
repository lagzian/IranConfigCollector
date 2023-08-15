import json
import base64
import pickle
import requests
from urllib.parse import urlparse

def decode_vmess(vmess_config):
    vmess_data = vmess_config[8:]  # Remove "vmess://"
    decoded_data = json.loads(base64.b64decode(vmess_data))
    return decoded_data

def encode_vmess(config):
    encoded_data = base64.b64encode(json.dumps(config).encode()).decode()
    vmess_config = "vmess://" + encoded_data
    return vmess_config

def remove_duplicate_vmess(input_data):
    array = input_data.split("\n")
    result = {}
    
    for item in array:
        parts = decode_vmess(item)
        if parts is not None and len(parts) >= 3:
            part_ps = parts["ps"]
            del parts["ps"]
            sorted_parts = dict(sorted(parts.items()))
            part_serialize = base64.b64encode(pickle.dumps(sorted_parts)).decode()
            result.setdefault(part_serialize, []).append(part_ps)

    final_result = []
    for serial, ps_list in result.items():
        part_after_hash = ps_list[0] if ps_list else ""
        part_serialize = pickle.loads(base64.b64decode(serial.encode()))
        part_serialize["ps"] = part_after_hash
        final_result.append(encode_vmess(part_serialize))
    
    output = "\n".join(final_result)
    return output

def main():
    url = "https://raw.githubusercontent.com/lagzian/TelegramV2rayCollector/main/sub/vmess_base64"
    response = requests.get(url)
    input_data = response.text
    
    cleaned_output = remove_duplicate_vmess(input_data)
    
    output_path = "./sub/splitted/"
    output_filename = "vmess_clean.txt"
    
    with open(output_path + output_filename, "w") as f:
        f.write(cleaned_output)
    
    print("Cleaned VMess configurations written to", output_path + output_filename)

if __name__ == "__main__":
    main()
