var rf = {

    //localStorage
    ls: function (k, v) {
        if (arguments.length == 2) {
            localStorage[k] = v;
        } else {
            return localStorage[k] || null;
        }
    },
   
    //Mobile
    isPC: function () { return !navigator.userAgent.match(/(iPhone|iPod|Android|ios)/i) },
    //DOM集合某一项是否包含节点
    contains: function (arr, target) {
        var item = null;
        $(arr).each(function () {
            if (this.contains(target)) {
                item = this;
                return false;
            }
        })
        return item;
    },
    //选项卡定位
    PositionTab: function () {
        var mm = $('#mtab-main'), vw = mm.parent().width(), sw = mm[0].scrollWidth;
        if (sw > vw) {
            var max = 0, min = -1 * (sw - vw), ml = 3, left = parseFloat(mm.css('left')), ca;
            mm.children().each(function () {
                ml += $(this).outerWidth() + 3;
                if (this.className.indexOf("active") >= 0) {
                    ca = this;
                    return false;
                }
            });
            if (ml + left > vw) {
                ml = -1 * ml + (vw / 2);
                ml = Math.max(min, ml);
                mm.css('left', ml);
            } else if (ml - $(ca).outerWidth() + left < 0) {
                ml = -1 * ml + (vw / 2);
                ml = Math.min(0, ml);
                mm.css('left', ml);
            }
 
        }
    },
    //打开选项卡 链接、含图标的标题，false不显示关闭按钮(可选)
    OpenPage: function (url, title, close) {
        var isopen = false, mmc = $('#mtab-main').children(), mb = $('#mtabox');
        mmc.removeClass('active');
        mb.children().removeClass('active');
        mb.find('iframe').each(function () {
            //var shorturl = (this.src.split(location.host)[1] || "").toLowerCase();
            //shorturl = shorturl.split('?')[0];

            //if (shorturl != "" && url.toLowerCase().indexOf(shorturl) > -1) {
            if($(this).data('id') == url){
                //取消注释，点击导航会刷新页面
                //this.src = url;
                var pageid = $(this).parent().addClass('active')[0].id;
                mmc.each(function () {
                    if (this.hash == "#" + pageid) {
                        $(this).addClass('active');
                        
                        return false;
                    }
                });
                
                isopen = true;
                return false;
            }
        });
        if (!isopen) {
            var pageid = "page_" + new Date().valueOf(), close = close == false ? '' : '<em class="fa fa-close"></em>';
            mmc.end().append('<a href="#' + pageid + '" class="active">' + title + close + '</a></li>');
            mb.append('<div class="tab-pane active" id="' + pageid + '"><iframe frameborder="0" src="about:blank" data-id="' + url + '"></iframe></div>');
            $('#' + pageid).find('iframe')[0].src = url;
        }

        //定位、调整
        rf.PositionTab();
        rf.IframeAuto();

        //Mobile 打开时隐藏菜单导航
        if (!rf.isPC()) {
            var mt = $('#menu-toggler');
            mt.hasClass('display') && $('#menu-toggler')[0].click();
        }
    },
    //Iframe自适应、延迟响应
    IframeAuto: function () {
        clearTimeout(window.deferIA);
        window.deferIA = setTimeout(function () {
            var footerhight = 0;

            if($('#footer').length > 0){   //如果footer不存在，则footerhight = 0
                if($('#footer').css("display") != 'none') {//如果footer隐藏，则footerhight = 0
                    footerhight = $('#footer').outerHeight();
                }
            }
            var box = $('#mtabox'), sh = $(window).height() - box.offset().top - footerhight; //foot内边距
            box.children('div').css("height", sh);
        }, 50);
    }
};

//关闭当前
function closeCurIframe() {
    var mmc = $('#mtab-main').children();
    mmc.each(function (i) {
        if(i) {
            if ($(this).hasClass('active')) {
                $(this).prev().addClass('active');
                $(this.hash).prev().addClass('active');

                $(this.hash).remove();
                $(this).remove();
                rf.PositionTab();
                return false;
            }
        }

    }).end().css('left', 0);
}



//点击选项卡
$('#mtab').click(function (e) {
    e = e || window.event;
    var target = e.target || e.srcElement;
    if (target.nodeName == "EM" && target.className.indexOf('fa-close') >= 0) {
        //关闭
        var ta = $(target).parent();
        if (ta.hasClass('active')) {
            ta.prev().addClass('active');
            $(ta[0].hash).prev().addClass('active');
        }
        $(ta[0].hash).remove();
        ta.remove();
        rf.PositionTab();
        return false;
    } else if ($(this).children('a')[0].contains(target)) {
        //左滑
        var mm = $('#mtab-main'), vw = mm.parent().width(),
            left = parseFloat(mm.css('left')) + (vw / 2);
        mm.css('left', Math.min(0, left));
    } else if ($(this).children('a').last()[0].contains(target)) {
        //右滑
        var mm = $('#mtab-main'), vw = mm.parent().width(),
            min = -1 * (mm[0].scrollWidth - mm.parent().width()),
            left = parseFloat(mm.css('left')) - (vw / 2);
        mm.css('left', Math.max(min, left));
    } else {
        var item = rf.contains($('#mtab-menu').children(), target);
        if (item) {
            
        } else {
            var mmc = $('#mtab-main').children();
            item = rf.contains(mmc, target);
            if (item) {
                if (item.className.indexOf('active') == -1) {
                    //选项卡标签
                    mmc.removeClass('active');
                    $(item).addClass("active");
                    $('#mtabox').children().removeClass('active');
                    $(item.hash).addClass('active');
                    rf.PositionTab();
                }

                if($('.mtab .dropdown').hasClass('open'))
                    $('.mtab .dropdown-toggle').dropdown('toggle');
                return false;
            }
        }
    }
});


$('.tabShowActive').click(function (e) {
    rf.PositionTab();
});

$('.tabCloseOther').click(function (e) {
    //关闭其他
    var mmc = $('#mtab-main').children();
    mmc.each(function (i) {
        if (i && this.className.indexOf("active") == -1) {
            $(this.hash).remove();
            $(this).remove();
        }
    }).end().css('left', 0);
});

$('.tabCloseAll').click(function (e) {
    //关闭全部
    var mmc = $('#mtab-main').children();
    if (mmc.length) {
        mmc.each(function (i) {
            if (i) {
                $(this.hash).remove();
                $(this).remove();
            }
         }).first().addClass('active');
         $(mmc.first()[0].hash).addClass('active');
         mmc.end().css('left', 0);
    }
});

$('.tabReFresh').click(function (e) {
     //刷新
    var currt = $('#mtabox').children('div.active').children();
    if (currt.length) {
        currt[0].contentWindow.location.reload(false);
        //currt[0].src = currt[0].src;
    }
});

function closeNav(){

    if ((navigator.userAgent.match(/(phone|pad|pod|iPhone|iPod|ios|iPad|Android|Mobile|BlackBerry|IEMobile|MQQBrowser|JUC|Fennec|wOSBrowser|BrowserNG|WebOS|Symbian|Windows Phone)/i))) {
        //alert('手机端')
        var container = $('body > #container');

        if(container.hasClass('push')){
            $.niftyNav('pushToggle');
        }else if(container.hasClass('slide')){
            $.niftyNav('slideToggle');
        }else if(container.hasClass('reveal')){
            $.niftyNav('revealToggle');
        }else{
            $.niftyNav('colExpToggle');
        }
    }else{
        //alert('PC端')
    }



}
//effect aside-float aside-bright page-fixedbar mainnav-fixed mainnav-in
//effect aside-float aside-bright page-fixedbar mainnav-fixed mainnav-sm lg-mainnav-lg
//effect aside-float aside-bright page-fixedbar mainnav-fixed lg-mainnav-lg mainnav-in
$('.menuItem').on('click',function (e) {
    e.preventDefault(); //阻止a的默认事件

    var url = $(this).attr("href"),
        name = this.innerHTML,
        label = $(this).find('.label').prop("outerHTML");
        badge = $(this).find('.badge').prop("outerHTML");
        if(label){
            name = name.replace(label,"").toString() 
        }
        if(badge){
            name = name.replace(badge,"").toString() 
        }

        rf.OpenPage(url, name, true);

        //延迟执行防止菜单栏隐藏问题
        setTimeout(function () {
            closeNav();
        },50);

});


$(function(){

   //rf.OpenPage("home.html", '<i class="fa fa-home"></i>首页', false);

});

// function(e) {
//     if (e && e.preventDefault) {
//       e.preventDefault()
//     } else {
//       window.event.returnValue = false
//     }
//     e = e || window.event;
//     var target = e.target || e.srcElement;
//     $(this).find('a').each(function() {
//       if (this.contains(target)) {
//         if (this.href.split('#')[1]) {
//           rf.OpenPage(this.href.split('#')[1], this.innerHTML);
//         }
//         return false;
//       }
//     });
//   }

$(window).resize(rf.IframeAuto);
