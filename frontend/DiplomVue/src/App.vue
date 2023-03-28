<template>
  <div class="border">
    <h1>Вход</h1>
    <input type="text" name="" id="" class="authorization_input" placeholder="Логин" v-model="login_input">
    <input type="text" name="" id="" class="authorization_input" placeholder="Пароль" v-model="password_input">
    <button id="border_button" @click="entrance()">Войти</button>
  </div>
</template>

<script>


export default {
  data() {
    return {
      login_input: ' ',
      password_input: ' '
    }
  },
  methods: {
    entrance() {
      let xhr = new XMLHttpRequest();
      xhr.open('POST', 'http://10.3.8.155:8001/auth');
      xhr.send(`{"uid":"${this.login_input.value}", "password": "${this.password_input.value}"}`);

      xhr.onload = function() {
        if (xhr.status !== 200) {
          console.log(xhr.status);
        } else {
          console.log('Yas');
        }
      }
    }
  }
}
</script>

<style>

  .border {
    border: 1px solid none;
    border-radius: 5%;
    width: 15%;
    padding: 20px;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    margin-left: auto;
    margin-right: auto;
    margin-top: 18%;
    box-shadow: 0 0 15px 2px #808080;
  }
  .authorization_input {
      margin-top: 25px;
      width: 100%;
      height: 30px;
      border-radius: 10px;
      outline: none;
  }
  #border_button {
    border-radius: 10px;
    margin-top: 30px;
    width: 100%;
    height: 40px;
    background-color: #427387;
    color: #ffff;
  }
  #border_button:hover {
    background-color: #2f5f72;
  }
</style>
