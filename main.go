name: Combined Workflow

on:
  schedule:
    - cron: '0 */4 * * *'
  workflow_dispatch:

jobs:
  sync_repository:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout repository
        uses: actions/checkout@v2

      - name: Install dependencies
        run: pip install python-telegram-bot

      - name: Sync repository
        env:
          GH_TOKEN: ${{ secrets.GH_TOKEN1 }}
        id: sync
        run: |
          gh repo sync lagzian/TelegramV2rayCollector -b main --force || echo "::set-output name=exit_code::1"

      - name: Send Telegram notification on success
        if: steps.sync.outputs.exit_code == '0'
        env:
          TELEGRAM_BOT_TOKEN: ${{ secrets.TELEGRAM_BOT_TOKEN }}
          TELEGRAM_CHAT_ID: ${{ secrets.TELEGRAM_CHAT_ID }}
        run: |
          python - <<EOF
          import asyncio
          from telegram import Bot

          async def send_telegram_message():
              bot_token = '${{ env.TELEGRAM_BOT_TOKEN }}'
              bot = Bot(token=bot_token)
              chat_id = '${{ env.TELEGRAM_CHAT_ID }}'
              message = '✅✅🔥 **V2RAYFIXER** V2ray Repository Updated Successfully! 🔥✅✅'
              await bot.send_message(chat_id=chat_id, text=message)

          asyncio.run(send_telegram_message())
          EOF

      - name: Send Telegram notification on failure
        if: steps.sync.outputs.exit_code != '0'
        env:
          TELEGRAM_BOT_TOKEN: ${{ secrets.TELEGRAM_BOT_TOKEN }}
          TELEGRAM_CHAT_ID: ${{ secrets.TELEGRAM_CHAT_ID }}
        run: |
          python - <<EOF
          import asyncio
          from telegram import Bot

          async def send_telegram_message():
              bot_token = '${{ env.TELEGRAM_BOT_TOKEN }}'
              bot = Bot(token=bot_token)
              chat_id = '${{ env.TELEGRAM_CHAT_ID }}'
              message = '❗❗⚠️⚠️🔥**V2RAYFIXER** V2ray Repository Sync Failed! 🔥⚠️⚠️❗❗'
              await bot.send_message(chat_id=chat_id, text=message)

          asyncio.run(send_telegram_message())
          EOF

  edit_readme:
    needs: sync_repository  # Ensure that the sync_repository job is completed before starting this job
    runs-on: ubuntu-latest

    steps:
      - name: Checkout repository
        uses: actions/checkout@v2
        with:
          repository: lagzian/TelegramV2rayCollector
          path: repo  # Clone the repository to the "repo" directory
          persist-credentials: false
          fetch-depth: 0

      - name: Replace text in README.md
        run: |
          sed -i 's/yebekhe/lagzian/g' repo/README.md
          cd repo
          git config user.name "lagzian"
          git config user.email "milad.lagzian@gmail.com"
          git add README.md

          # Check if there are any changes to commit
          if git diff-index --quiet HEAD; then
            echo "No changes to commit."
          else
            git commit -m "Replace 'yebekhe' with 'lagzian'"
            git remote set-url origin https://${{ secrets.ACCESS_TOKEN }}@github.com/lagzian/TelegramV2rayCollector.git
            git push origin HEAD:main
          fi
