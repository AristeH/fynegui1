# fynegui
Приложение рисует формы на основание Метаданных.
После авторизации генерируется основная форма приложения. Данные берутся с сервера.
Таблица Подсистемы-горизонтальная строка комманд.
имя файла md 
При загрузке Подсистемы загружаем Id таблиц в pudge
    сохраняем id подсистемы и Список таблиц id
    гуи_ID список параметров
При загрузке таблиц в pudge пишем 
    id  таблицы и список реквизитов
    влад_id  список владельцев таблиц id
    подч_id список подчиненных таблиц id
    гуи_id список параметров   
При загрузке реквизитов
    гуи_id список параметров

список параметров
    namerus
    nameeng
    synonym
    por
    mesto_v {top, bottom, left, right}
    type_v  {form, list, accord, table, tree, tool}
    name_type_v 
    width

имя файла данные
id таблица данных(id_рек_гуи значение )
