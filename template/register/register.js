function Jump(){
    var user_name  = $("#user_name").val()
    var password  = $("#password").val()
    var str = window.sessionStorage;
    $.ajax({
        url: str.getItem("domain_name")+"/mini/user/register",
        type: 'POST',
        contentType: "application/json",
        data:JSON.stringify({user_name:user_name,pass_word:password}),
        success: function (res) {
            console.log(res)
            if (res.message=="SUCCESS"){
                window.location.href='../login/login.html';
            }else {
                alert(res.message)
            }
        },
        error: function (res) {
            console.log(res)
            alert(res.message)
        }
    })

}

