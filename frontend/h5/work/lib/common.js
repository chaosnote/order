function getCommonButton(ref) {
    var menu = [
        {
            label:"本日店家",
            url:"/user/shop_list.html"
        },
        {
            label:"購物車",
            url:"/user/shop_car.html"
        }
    ] ;
    for (var key in ref){
        menu.push({
            label : ref[key], 
            url : key
        }) ;
    }
    return menu ;
}