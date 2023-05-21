<template>
  <div class="container d-flex justify-content-center align-items-center vh-100">
    <div class="card p-4">
      <h3 class="card-title text-center mb-4">Авторизация</h3>
      <form @submit.prevent="login">
        <div class="form-group m-4">
          <label for="username">Логин</label>
          <input type="text" id="username" class="form-control" v-model="uid" required>
        </div>
        <div class="form-group m-4">
          <label for="password">Пароль</label>
          <input type="password" id="password" class="form-control" v-model="password" required>
        </div>
        <div class="text-center m-4">
          <button type="submit" class="btn btn-primary">Войти</button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  name: "AuthPage",
  data() {
    return {
      uid: 'i21s617',
      password: 'ShBYjnsj'
    };
  },
  methods: {
    login() {

      const self = this; // Сохраняем ссылку на объект this

      const data = JSON.stringify({
        uid: this.uid,
        password: this.password
      });

      const xhr = new XMLHttpRequest();
      xhr.withCredentials = true;

      xhr.addEventListener("readystatechange", function() {
        if (this.readyState === 4) {
          if (this.status === 200) {
            const response = JSON.parse(this.responseText);
            if (response.valid === true) {
              // Перенаправление на MainPage
              self.$router.push("/main"); // Используем сохраненную ссылку self
            }
          }
        }
      });

      xhr.open("POST", "http://localhost:8001/auth/auth");
      xhr.setRequestHeader("Content-Type", "application/json");

      xhr.send(data);
    }
  }
}
</script>

<style scoped>
.container {
}

.card {
  max-width: 400px;
  width: 100%;
  padding: 20px;
}
</style>