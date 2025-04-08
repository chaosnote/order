function validateUserName(username) {
    const regex = /^[a-zA-Z][a-zA-Z0-9]{3,7}$/;
    return regex.test(username);
}

function validatePassword(password) {
    const regex = /^[a-zA-Z0-9]{3,7}$/;
    return regex.test(password);
}

function validateNickname(username) {
    const regex = /^[a-zA-Z\u4E00-\u9FFF][a-zA-Z0-9\u4E00-\u9FFF]{1,10}$/;
    return regex.test(username);
}

function validateShopName(name) {
    const regex = /^[a-zA-Z0-9\u4E00-\u9FFF]{1,10}$/;
    return regex.test(name);
}

function validateShopMobile(mobile) {
    const regex = /^[0-9]{9,10}$/;
    return regex.test(mobile);
}

function validateItemName(name) {
    const regex = /^[a-zA-Z0-9\u4E00-\u9FFF|-]{1,10}$/;
    return regex.test(name);
}

function validateItemPrice(price) {
    const regex = /^[1-9][0-9]{0,4}$/;
    return regex.test(price);
}
