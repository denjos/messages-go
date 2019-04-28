let formComment = $('#form-comment-add'),
    commentMessage = $('#comment-message'),
    commentContent = $('#comment-content');

formComment.addEventListener('submit', (el) => {
    el.preventDefault();
    let comment = {
        content: commentContent.value
    };
    requestAjax(formComment.method, formComment.action, JSON.stringify(comment))
        .then((rpt) => {
            if (rpt.status === 201) {
                commentMessage.textContent = rpt.response.message;
                // render comment
            } else {
                commentMessage.textContent = rpt.response.message;
            }
        }).catch((error) => {
        console.log(error);
    });
});


function getComment() {
    requestAjax('GET','/api/comments').then((rpt)=>{
        let msg=rpt.response;

        msg.forEach((el)=>{
            let div=document.createElement('div');
            div.innerText=el.content;
            container.appendChild(div);
        });
        //console.log(rpt.response);
    }).catch((error)=>{
        console.log(error);
    })
}

getComment();