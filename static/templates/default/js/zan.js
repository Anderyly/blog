/**
 * ZanBlog JavaScript File
 *
 * 为了提高用户使用ZanBlog时的用户体验。
 *
 * Author: 佚站互联（YEAHZAN）
 *
 * Site: http://www.yeahzan.com/
 */

jQuery(window).load(function() {
  zan.flexslider();
});


jQuery(function($) {
	zan.init();
});

/*audiojs.events.ready(function() {
    var as = audiojs.createAll();
});*/

var zan = {

	//初始化函数
	init: function() {
		this.dropDown();
		this.setImgHeight();
    this.addAnimation();
    this.archivesNum();
    this.scrollTop();
    this.ajaxCommentsPage();
	},

	// 设置导航栏子菜单下拉交互
	dropDown: function() {
		var dropDownLi = jQuery('.nav.navbar-nav li');

		dropDownLi.mouseover(function() {
			jQuery(this).addClass('open');
		}).mouseout(function() {
			jQuery(this).removeClass('open');
		});
	},

	// 设置文章图片高度
	setImgHeight: function() {
		var img = jQuery(".zan-single-content").find("img");

		img.each(function() {
			var $this 		 = jQuery(this),
					attrWidth  = $this.attr('width'),
					attrHeight = $this.attr('height'),
					width 		 = $this.width(),
					scale      = width / attrWidth,
					height     = scale * attrHeight;

			$this.css('height', height);

		});
	},

  // 为指定元素添加动态样式
  addAnimation: function() {
    var animations = jQuery("[data-toggle='animation']");

    animations.each(function() {
      jQuery(this).addClass("animation", 2000);
    });
  },

	// 设置首页幻灯片
	flexslider: function() {
		jQuery('.flexslider').flexslider({
	    animation: "slide"
	  });
	},

	// 设置每月文章数量
	 archivesNum: function() {
		jQuery('#archives .panel-body').each(function() {
			var num = jQuery(this).find('p').size();
			var archiveA = jQuery(this).parent().siblings().find("a");
			var text = archiveA.text();

			archiveA.html(text + ' <small>(' + num + '篇文章)</small>');
		});
 	},

 	//头部固定
 	scrollTop: function() {
		//获取要定位元素距离浏览器顶部的距离 
		var navH = jQuery("#zan-nav").offset().top; 

		//滚动条事件 
		jQuery(window).scroll(function(){ 
			//获取滚动条的滑动距离 
			var scroH = jQuery(this).scrollTop(); 

			//滚动条的滑动距离大于等于定位元素距离浏览器顶部的距离，就固定，反之就不固定 
			if(scroH>=navH){ 
				//jQuery("#zan-nav").addClass("navbar-fixed-top"); 

			}else if(scroH<navH){ 
				jQuery("#zan-nav").removeClass("navbar-fixed-top"); 
			} 

		}); 
	},

	// ajax评论分页
	ajaxCommentsPage: function() {
		//var $ = jQuery.noConflict();

		$body=(window.opera)?(document.compatMode=="CSS1Compat"?$('html'):$('body')):$('html,body');
		// 点击分页导航链接时触发分页
		$('#comment-nav a').live('click', function(e) {
		    e.preventDefault();
		    $.ajax({
		        type: "GET",
		        url: $(this).attr('href'),
		        beforeSend: function(){
		            $('#comment-nav').remove();
		            $('#comment-list').remove();
		            $('#loading-comments').slideDown();
		            //$body.animate({scrollTop: $('#comments-number').offset().top - 65}, 800 );
		        },
		        dataType: "html",
		        success: function(out){
		            var result = $(out).find('#comment-list');
		            var nextlink = $(out).find('#comment-nav');

		            $('#loading-comments').slideUp('fast');
		            $('#loading-comments').after(result);
		            $('#comment-list').after(nextlink);
		        },
				error: function(error) {
					alert('创建连接失败');
				}
		    });
		});
	}
}

/* jQuery Ajax 提交、刷新 评论   本插件由 简爱 http://www.gouji.org/ 提取完善、移植 至 EMLOG */

// JS 仿 PHP 获取 GET 数据
var $_GET=(function(){
  var js=document.scripts;
  var url=js[js.length-1].src;
  var u=url.split("?");
  if(typeof(u[1])=="string"){
    u=u[1].split("&");
    var get={};
    for(var i in u){
      var j=u[i].split("=");
      get[j[0]]=decodeURIComponent(j[1]);
    }
    return get;
  }
  else{
    return {};
  }
})();


$(function(){
  JA_comment_($_GET['list'], $_GET['form'], $_GET['css'], $_GET['msg']);
});


function JA_comment_(JA_commentList, JA_commentForm, JA_commentCss, JA_comment) {
    if (!JA_commentList) var JA_commentList = "#comment-list";
    if (!JA_commentForm) var JA_commentForm = "#commentform";
    var JA_commentFormSubmit = $(JA_commentForm).find("input[type=submit]");
    var JA_commentFormT = $(JA_commentForm + ' textarea');
    var JA_commentForm = $(JA_commentForm);
    JA_commentForm.submit(function () {
        var q = JA_commentForm.serialize();
        JA_commentFormT.attr("disabled", true);
        JA_commentFormSubmit.attr("disabled", true);
        $("#JA_commenError").text('');
		$("#JA_commenSuccess").text('');
        $("#JA_commenLoading").show();
        $.post(JA_commentForm.attr("action"), q, function (d) {
            var reg = /<div class=\"main\">[\r\n]*<p>(.*?)<\/p>/i;
            if (reg.test(d)) {
				if(d.indexOf("评论发表成功")>0)
					$("#JA_commenSuccess").html(d.match(reg)[1]);
				else
					$("#JA_commenError").html(d.match(reg)[1]);
                $("#JA_commenLoading").hide()
            } else {
                var p = $("input[name=pid]").val();
                cancelReply();
                $("[name=comment]").val("");
                $("#JA_commenError").text('');
                u = JA_commentList.split(",");
                for (var i in u) {
                    $(u[i]).html($(d).find(u[i]).html())
                };
                var body = (window.opera) ? (document.compatMode == "CSS1Compat" ? $('html') : $('body')) : $('html,body');
                if (p != 0) {
                    var body = (window.opera) ? (document.compatMode == "CSS1Compat" ? $('html') : $('body')) : $('html,body');
                    body.animate({
                        scrollTop: $("#comment-" + p).offset().top - 20
                    }, "normal", function () {
                        $("#JA_commenLoading").hide()
                    })
                } else {
                    var body = (window.opera) ? (document.compatMode == "CSS1Compat" ? $('html') : $('body')) : $('html,body');
                    body.animate({
                        scrollTop: $(JA_commentList).offset().top - 20
                    }, "normal", function () {
                        $("#JA_commenLoading").hide()
                    })
                }
            };
            JA_commentFormT.attr("disabled", false);
            JA_commentFormSubmit.attr("disabled", false)
        });
        return false
    });
    JA_commentForm.keypress(function (e) {
        if (e.ctrlKey && e.which == 13 || e.which == 10) {
            JA_commentForm.submit()
        } else if (e.shiftKey && e.which == 13 || e.which == 10) {
            JA_commentForm.submit()
        }
    });
    msg = ' jQuery Ajax 提交、刷新 评论';
    JA_comment = JA_comment ? JA_comment : '评论提交中... 请稍候';
    if (JA_commentCss) {
        css = '<link type="text/css" rel="stylesheet" href="' + JA_commentCss + '" /><!--' + msg + '-->'
    } else {
        css = '<style type="text/css">/*' + msg + '*/#JA_commenLoading{font-size:12px;background:url(data:images/gif;base64,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) no-repeat left center;padding-left:22px;margin-bottom:5px}#JA_commenError{font-size:12px;background:url(data:images/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAABdElEQVQ4y6WTS0sCcRTFTfC5UEyZhYj4xAcoiowwDIKSQRtxkQgK4i5wFUGua6frdkKupXVfIMMv0KZFtVCLTAgFWxXp6c4EE4OPtBaHyz338PvfyzAKAIqlqlR2Vs5Jy4eNhg8+X48gu38DxOOXUCqBUOhmc0C1moTJNIHVCuj1M9qivAlACafzCpHIJ4rFA6pj2O335OvWA+Tz+zAYPsCyT9RrwHG30OmmyGaPfwcMhxaEwx3xdpZ9pVPsiMUexN7lukO77V4NSKVKsFi+7WBQeNUBr/dF7I1G0DYnlNtaDGi1thGNdkVLkM0GFAohMMxI8vz+Eer1wGJAJnMIjQZSWKsFkslr8qaSp1IBPH9GeZUcUKuZaNV3KSjI7Z7Q/afweAYy3+EASqWAHJDLnctCglh2jGZzj+7uzs3S6Qs5gGEwF1KrZ/Q536hO52ZmM1AuH/0AOO4ZiUSP1JfE833yB1QfZb6QE/K5XGf1v7Cm/g34AiZUmQgJdHT+AAAAAElFTkSuQmCC) no-repeat left center;padding-left:22px;margin-top:5px;color:red}#JA_commenSuccess{font-size:12px;background:url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAgY0hSTQAAeiYAAICEAAD6AAAAgOgAAHUwAADqYAAAOpgAABdwnLpRPAAAAAlwSFlzAAAk6AAAJOgBgmMFHAAAAixJREFUOE99U71rFEEUn0MbTQLZeXOIIab0bndn97C2ENIlgvkLxP/ASgtbLeyDIkIKFREMXHb24y6axJwQ0UYUC8WIhZWCYKGFaIwZf29yt94lpwfDzb6Z9/t4701FDPlN3p88ND7yY9ru2JM4HrMH7Ob278rauzNf3g67X8ZqCzQW5HQlbNGHaFlZ3cZqkY0eKBvmtIW1GqQ0PRQkTLxjYUEbjU51NyEjG6Z/ly7Ixo+UjZbppzby/ABIsFgdBfLT+HF1X2IJAkAG0VAWPQRBQmdLELBdjjsI9jEO7HMkwkqQyu+6jb1TKD8HmTcl4jtHRkIjP7ogAzATEkoA7Fl6mMgbsVEnAPKcVcTriBm6KPxczXLBeoksEwc7zMgxl2zoFuRWWHKYyrvRCmysunhb+Im8yrJQYYvqb+lUXtJGndaZ/NbYqEI23S6TM5rvFZhJA0OvGOAas+kWEAv1tZaMR8zkpzSnc3ld2C6zofl4DW3t2mMbAHgjfLTEocI/B8NCbtabnu5vE6wMJPNd1wlDz0R9SR1H4JfzzgfMUsj3OqO68zwkme9xbaBgwRGhsp1usXZBuEApvUZ37rmJ7O8KdwrfrMBf8mYcgM7UKRRw21nozoIbGFR770Q6dh44I43oiIOlVVcLoDpv/xmoBicX8kXN0MS+NwEr53RLfuIhYRu9B8Wg7h3gHx1r1puHj/7zVQaL3hQe1QWoaEP+S24VXugT+L7pJ2p2b+Ifh4ZuACm9XXoAAAAASUVORK5CYII=) no-repeat left center;padding-left:22px;margin-top:5px;color:green}</style>'
    };
    JA_commentFormSubmit.before(css + '<div id="JA_commenLoading" style="display:none;">' + JA_comment + '</div><div id="JA_commenError"></div><div id="JA_commenSuccess"></div>')
}