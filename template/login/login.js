function login(){
    var user_name  = $("#user_name").val()
    var password  = $("#password").val()
    var str = window.sessionStorage;
    $.ajax({
        url: str.getItem("domain_name")+"/mini/user/login/",
        type: 'POST',
        contentType: "application/json",
        data:JSON.stringify({user_name:user_name,pass_word:password}),
        success: function (res) {
            console.log(res)
            console.log(res.token)
            var storage = window.sessionStorage;
            storage.setItem('token', res.token);
            window.location.href='../index/index.html';
            // window.location.href="http://www.baidu.com"; //在原有窗口打开
        },
        error: function (res) {
            alert(res)
        }
    })
}

