function InsertMemo(){
    var tit  = $("#tit").val();
    var con  = $("#con").val();
    var str = window.sessionStorage;
    $.ajax({
        headers:{
            "token":str.getItem("token"),
        },
        url: str.getItem("domain_name")+"/miniProgram/insert/one/content",
        type: 'POST',
        contentType: "application/json",
        data:JSON.stringify({title:tit,content:con}),
        success: function (res) {
            console.log(res);
            if (res.message=="SUCCESS"){
                window.location.href='../index/index.html';
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