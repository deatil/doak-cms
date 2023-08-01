(function($) {
    "use strict";
    
    // 退出
    $(".js-logout-btn").click(function(e) {
        e.stopPropagation;
        e.preventDefault;

        var url = $(this).data("url");
        layer.confirm('您确定要退出登陆吗？', { 
            icon: 3, 
            title: '提示信息' 
        }, function(index) {
            location.href = url;
        });
    });
})(jQuery);