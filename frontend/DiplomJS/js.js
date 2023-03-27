function authorization() {
let login = document.getElementById('login');
let password = document.getElementById('password');
let xhr = new XMLHttpRequest();
xhr.open('POST', 'http://10.3.8.155:8001/auth');
xhr.send(`{"uid":"${login.value}", "password":"${password.value}"}`);

xhr.onload = function() {
  if (xhr.status !== 200) {
    console.log('Пошла нахуй');
  } else {
    document.location.href = "index2.html";
    }
  }
}
