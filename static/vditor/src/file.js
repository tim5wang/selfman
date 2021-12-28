
var getParams = function () {
    var url = window.location.search
    var queryParamsString = url.substring(1, url.length)
    var queryParams = queryParamsString.split('&')
    var params = {}
    if (queryParamsString.length) {
        queryParams.forEach(function (queryParam) {
            var splittedParam = queryParam.split('=')
            var param = splittedParam[0]
            var value = splittedParam[1]
            params[param] = value
        })
    }
    return params
}

