let email = $('#email'),
    pwd = $('#password'),
    form = $('#formLogin'),
    message = $('#message-login');
form.addEventListener('submit', (e) => {
    e.preventDefault();
    let request = {
        email: email.value,
        password: pwd.value
    };
    requestAjax(form.method, form.action, JSON.stringify(request)).then((response) => {

        if (response.status === 200) {
            message.textContent = 'login es successful';
            let token = response.response.token;
            sessionStorage.setItem('token', token);
            console.log(response.response);
        } else {
            console.log(response.response);
            message.textContent = response.response.message;
        }

    }).catch((error) => {
        console.log(error);
    });
});

