function login(){
    var user_name  = $("#user_name").val();
    var password  = $("#password").val();
    var str = window.sessionStorage;
    $.ajax({
        url: str.getItem("domain_name")+"/mini/user/login/",
        type: 'POST',
        contentType: "application/json",
        data:JSON.stringify({user_name:user_name,pass_word:password}),
        success: function (res) {
            console.log(res);
            console.log(res.token);
            if (res.message=="SUCCESS"){
                var storage = window.sessionStorage;
                storage.setItem('token', res.token);
                storage.setItem('user_name', user_name);
                window.location.href='../hl/index.html';
            }else {
                alert(res.message)
            }
            // window.location.href="http://www.baidu.com"; //在原有窗口打开
        },
        error: function (res) {
            alert(res.message)
        }
    });
}

function Register(){
    window.location.href='../hl/register.html';
}