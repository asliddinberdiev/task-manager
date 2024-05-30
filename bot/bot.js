const TelegramBot = require('node-telegram-bot-api');

const token = process.env.BOT_TOKEN ?? "7112040080:AAEH5tKV124Q3JveP_NWmro2QxU-ZFfojc4";

const bot = new TelegramBot(token, {polling: true});

bot.on('message', (msg) => {
  const chatId = msg.chat.id;

  bot.sendMessage(chatId, "test");
});