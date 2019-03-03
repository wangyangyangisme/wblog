(function() {
    var obj = {};
    obj.loadScript = function(url, callback) {
        var doc = document;
        var script = doc.createElement("script");
        script.type = "text/javascript";
        if (script.readyState) {
            script.onreadystatechange = function() {
                if (script.readyState == "load" || script.readyState == "complete") {
                    script.onreadystatechange = null;
                    callback()
                }
            }
        } else {
            script.onload = function() {
                callback()
            }
        }
        script.src = url;
        doc.getElementsByTagName("head")[0].appendChild(script)
    };
    var jsList = [qiniu_domain + "/static/js/jquery.min.js", qiniu_domain + "/static/dux/js/libs/bootstrap.min.js", qiniu_domain + "/static/dux/js/loader.js"];
    function callback() {
        jsList.length ? obj.loadScript(jsList.shift(), callback) : (function() {
            time = null
        })()
    }
    var time = setTimeout(function() {
            obj.loadScript(jsList.shift(), callback)
        },
        100);
    setTimeout(function() {
            tbquire(["page"],
                function() {
                    $("#page").jqPaginator({
                        totalPages: TotalPage,
                        visiblePages: 1,
                        first: '<li class="first"><a href="javascript:;">首 页</a></li>',
                        prev: '<li class="prev"><a href="javascript:;">上一页</a></li>',
                        next: '<li class="next"><a href="javascript:;">下一页</a></li>',
                        last: '<li class="last"><a href="javascript:;">尾 页</a></li>',
                        currentPage: currentPage,
                        onPageChange: function(num, type) {
                            if (type == "change") {
                                window.location.href = "/article_cate?cid=" + cid + "&page=" + num
                            }
                        }
                    })
                })
        },
        1000)
})();