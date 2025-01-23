// ОБОРАЧИВАЮ В ASYNC, ЧТОБЫ FETCH-ФУНКЦИИ ВЫПОЛНЯЛИСЬ ПОСЛЕДОВАТЕЛЬНО
(async () => {
    const URL = 'https://examly.ru/api/v1'


    // РЕГИСТРАЦИЯ ПОЛЬЗОВАТЕЛЯ
    //
    // fetch(`${URL}/auth/register`, {
    //     method: 'POST',
    //     body: JSON.stringify({
    //         role: 'member',
    //         email: 'ivan@mail.ru',
    //         first_name: 'Ivan',
    //         last_name: 'Ivanov',
    //         password: 'password',
    //         password_confirmation: 'password',
    //     }),
    // }).then(data => data.json()).then(data => console.log(data))


    let token = ''

    // ПОЛУЧЕНИЕ JWT-ТОКЕНА (НУЖЕН ДЛЯ ВСЕХ БУДУЩИХ ЗАПРОСОВ)
    // ДАННЫЕ ДЛЯ АДМИНОК СКИНУЛ В МЕССЕНДЖЕР
    await fetch(`${URL}/auth/login`, {
        method: 'POST',
        body: JSON.stringify({
            email: 'ivan@mail.ru',
            password: 'password',
        }),
    }).then(data => data.json()).then(data => token = data.data)


    let uuid = '01944122-9752-7c36-b4b0-06689fa0a502'

    // СОЗДАНИЕ КУРСА
    //
    // СПИСОК ЦВЕТОВ (color): https://github.com/zagvozdeen/examly/blob/main/resources/ts/components/AppColorSelector.vue
    // СПИСОК ИКОНОК (icon): https://github.com/zagvozdeen/examly/blob/main/resources/ts/composables/useIcons.ts
    //
    // await fetch(`${URL}/courses`, {
    //     method: 'POST',
    //     headers: {
    //         Authorization: `Bearer ${token}`,
    //     },
    //     body: JSON.stringify({
    //         name: 'Вопросы с собеседования на Backend Golang Developer',
    //         description: 'В этом курсе собраны все вопросы с собеседований на Backend Golang Developer',
    //         color: 'bg-emerald-400',
    //         icon: 'airplane-fill',
    //     }),
    // }).then(data => data.json()).then(data => console.log(data))


    // КАЖДАЯ СУЩНОСТЬ ИЗ СПИСКА (course, module, question) ДОЛЖЕН ПРОЙТИ МОДЕРАЦИЮ,
    // ПОЭТОМУ ПОСЛЕ СОЗДАНИЯ СУЩНОСТИ МЫ ПРИСВАИВАЕМ ЕЙ АКТИВНЫЙ СТАТУС
    //
    // СПИСОК СТАТУСОВ (status): https://github.com/zagvozdeen/examly/blob/main/internal/enum/status.go
    // НАМ ВЕЗДЕ НУЖЕН СТАТУС «active»
    //
    // ВАЖНО! МОДЕРИРОВАТЬ МОЖНО ТОЛЬКО С АДМИНСКОГО АККАУНТА, ДАННЫЕ ОТ НИХ СКИНУЛ В МЕССЕНДЖЕР
    // await fetch(`${URL}/courses/${uuid}/moderate`, {
    //     method: 'PATCH',
    //     headers: {
    //         Authorization: `Bearer ${token}`,
    //     },
    //     body: JSON.stringify({
    //         status: 'active',
    //     })
    // }).then(data => data.json()).then(data => console.log(data))


    uuid = '01944130-9c96-77a4-b5c9-5ff1a5ae7d28'

    // СОЗДАНИЕ МОДУЛЯ
    //
    // await fetch(`${URL}/modules`, {
    //     method: 'POST',
    //     headers: {
    //         Authorization: `Bearer ${token}`,
    //     },
    //     body: JSON.stringify({
    //         name: 'Модуль «Backend»',
    //         course_id: 1,
    //     }),
    // }).then(data => data.json()).then(data => console.log(data))


    // МОДЕРАЦИЯ МОДУЛЯ
    // await fetch(`${URL}/modules/${uuid}/moderate`, {
    //     method: 'PATCH',
    //     headers: {
    //         Authorization: `Bearer ${token}`,
    //     },
    //     body: JSON.stringify({
    //         status: 'active',
    //     })
    // }).then(data => data.json()).then(data => console.log(data))


    uuid = '0194413a-1667-79a9-93dc-b73aa3a28ec9'

    // СОЗДАНИЕ ВОПРОСА
    //
    // СПИСОК ТИПОВ ВОПРОСОВ (type): https://github.com/zagvozdeen/examly/blob/main/internal/enum/question_type.go
    // СПИСОК ТЕГОВ И ИХ ID 136-174 (tags_ids): https://github.com/zagvozdeen/examly/blob/third-sprint-recommendations-and-personal-account-for-companies/migrations/000001_create_tables.up.sql
    // await fetch(`${URL}/questions`, {
    //     method: 'POST',
    //     headers: {
    //         Authorization: `Bearer ${token}`,
    //     },
    //     body: JSON.stringify({
    //         title: 'Что такое Golang?',
    //         content: 'Дайте краткий ответ', // необязательный параметр
    //         explanation: 'Golang — это язык программирования', // необязательный параметр
    //         type: 'single_choice',
    //         course_id: 1,
    //         module_id: 1, // необязательный параметр
    //         options: [{
    //             id: 1,
    //             content: 'Golang – это малоизвестная операционная система, разработанная для работы с микроконтроллерами.',
    //             is_correct: false,
    //         }, {
    //             id: 2,
    //             content: 'Golang – это язык программирования, разработанный в Google, известный своей простотой и эффективностью для создания масштабируемых серверных приложений.',
    //             is_correct: true,
    //         }, {
    //             id: 3,
    //             content: 'Golang – это графический редактор, популярный среди дизайнеров благодаря своим мощным инструментам для 3D-моделирования.',
    //             is_correct: false,
    //         }, {
    //             id: 4,
    //             content: 'Golang – это система управления базами данных, которая используется для анализа больших объемов данных.',
    //             is_correct: false,
    //         }],
    //         tags_ids: [1, 3, 35],
    //     }),
    // }).then(data => data.json()).then(data => console.log(data))


    // МОДЕРАЦИЯ ВОПРОСА
    // await fetch(`${URL}/questions/${uuid}/moderate`, {
    //     method: 'PATCH',
    //     headers: {
    //         Authorization: `Bearer ${token}`,
    //     },
    //     body: JSON.stringify({
    //         status: 'active',
    //     })
    // }).then(data => data.json()).then(data => console.log(data))


    // ЗАПОЛНИТЬ АНКЕТУ ОБРАТНОЙ СВЯЗИ (ПРИГОДИТСЯ В ПРЕЗЕНТАЦИИ)
    // await fetch(`${URL}/users/experience`, {
    //     method: 'POST',
    //     headers: {
    //         Authorization: `Bearer ${token}`,
    //     },
    //     body: JSON.stringify({
    //         one: 5, // от 1 до 5
    //         two: 9, // от 1 до 10
    //         three: 5, // от 1 до 5
    //         four: 'Раздел с советами по подготовке к техническим собеседованиям был особенно полезен, так как он содержал практические примеры и типичные вопросы.',
    //         five: 5, // от 1 до 5
    //         six: 6, // СВОБОДНЫЙ ВВОД ЧИСЛА (СКОЛЬКО ПРЕДЛОЖЕНИЙ О РАБОТЕ ПОЛУЧИЛ ЧЕЛОВЕК)
    //         seven: 'На данный момент я считаю, что все функции работают отлично.',
    //         eight: 'Нет, я не испытывал никаких проблем с навигацией и использованием ресурса.',
    //         nine: 5, // от 1 до 5
    //         ten: 'Может быть, больше информации о soft skills и их важности во время собеседований.',
    //         eleven: 5, // от 1 до 5
    //         twelve: 'В данный момент не вижу необходимости в изменениях, всё устраивает.',
    //         thirteen: 'Да, я смог получить работу благодаря вашему ресурсу. Процесс включал изучение предоставленных материалов, подачу резюме через ресурс и успешное прохождение собеседования с подготовкой по вашим методикам.',
    //     })
    // }).then(data => data.json()).then(data => console.log(data))
})()