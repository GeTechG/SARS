<template>
  <div class="container d-flex justify-content-center align-items-center vh-100">
    <div class="card p-4">
      <h3 class="card-title text-center mb-4">Авторизация</h3>
      <div v-if="showError" class="alert alert-danger text-center" role="alert">
        Данные для авторизации недействительны
      </div>
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
import { API_HOST } from "@/config";

export default {
  name: "AuthPage",
  data() {
    return {
      uid: '',
      password: '',
      showError: false
    };
  },
  created() {
    this.checkAuthentication();
  },
  methods: {
    checkAuthentication() {
      const self = this;

      const xhr = new XMLHttpRequest();
      xhr.withCredentials = true;

      xhr.addEventListener("readystatechange", function() {
        if (this.readyState === 4) {
          if (this.status === 200) {
            // Пользователь уже авторизован, перенаправляем на MainPage
            self.$router.push("/main");
          }
        } else if (this.status === 401) {
          // Пользователь не авторизован
          self.showError = false; // Скрываем сообщение об ошибке
        }
      });

      xhr.open("GET", `${API_HOST}/auth/is_auth`);
      xhr.send();
    },
    login() {
      const self = this;
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
              self.showError = false; // Очищаем флаг ошибки
              self.$router.push("/main");
            } else {
              self.showError = true; // Устанавливаем флаг ошибки
            }
          } else if (this.status === 406 || this.status === 500) {
            self.showError = true; // Устанавливаем флаг ошибки
          }
        }
      });

      xhr.open("POST", `${API_HOST}/auth/auth`);
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
