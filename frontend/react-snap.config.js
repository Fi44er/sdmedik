module.exports = {
  // Укажите пути, которые нужно предварительно рендерить
  include: [
    "/",
    "/about",
    "/products/*",
    "/catalog",
    "/delivery",
    "/returnpolicy",
    "/deteils",
    "/certificate",
    "/contacts",
  ], // Пример: рендеринг главной страницы, страницы "О нас" и всех страниц продуктов
  // Укажите пути, которые нужно игнорировать
  exclude: [
    "/auth",
    "/register",
    "/paymants",
    "/paymants/:id",
    "/profile",
    `/admin/*`,
  ], // Пример: игнорирование страниц входа и регистрации
  // Укажите дополнительные параметры
  puppeteer: {
    headless: true, // Запуск в безголовом режиме
  },
};
