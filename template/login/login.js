function login(){
    var user_name  = $("#user_name").val()
    var password  = $("#password").val()
    var str = window.sessionStorage;
    $.ajax({
        headers: {
            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX25hbWUiOiLmlrDkurpfMDA4NTE2SkoiLCJleHAiOjE2MDAxMzA3Nzd9.gtOOxV2JlOkRHzbeYBwMWq_tVDWCSuMQetwB13l6fis"
        },
        url: str.getItem("url")+"/mini/user/login/",
        type: 'POST',
        contentType: "application/json",
        data:JSON.stringify({user_name:user_name,pass_word:password}),
        success: function (res) {
            console.log(res)
            console.log(res.token)
            var storage = window.sessionStorage;
            storage.setItem('token', res.token);
            window.location.href='index.html';
            // window.location.href="http://www.baidu.com"; //在原有窗口打开
        },
        error: function (res) {
            console.log(res)
        }
    })

}

