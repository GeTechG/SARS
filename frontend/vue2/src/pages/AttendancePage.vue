<template>
  <div>
    <h1>Посещаемость студентов пары</h1>

    <table class="table">
      <thead>
      <tr>
        <th>Студент</th>
        <th>Посещение</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="member in members" :key="member.uid">
        <td>{{ member.cn }}</td>
        <td>
          <div class="form-check">
            <input class="form-check-input" type="checkbox" v-model="attendance[member.uid]" id="attendanceCheck" />
            <label class="form-check-label" for="attendanceCheck">{{ attendance[member.uid] ? 'Был' : 'Не был' }}</label>
          </div>
        </td>
      </tr>
      </tbody>
    </table>

    <div class="text-center mt-3">
      <button class="btn btn-primary mx-2" @click="saveAttendance">Сохранить</button>
      <button class="btn btn-secondary mx-2" @click="goToMainPage">Вернуться на главную страницу</button>
    </div>
  </div>
</template>

<script>
import {API_HOST} from "@/config";

export default {
  name: 'AttendancePage',
  data() {
    return {
      classes: [],
      members: [],
      attendance: {}
    };
  },
  async created() {
    this.checkAuthentication();
    await this.getClasses();
    let classId = this.$route.params.classId;
    let clazz = this.classes.find(value => Number(value.id) === Number(classId))
    if (clazz) {
      this.getMembers(clazz.group)
      this.getAttendance(clazz.id);
    }
  },
  methods: {
    checkAuthentication() {
      const self = this;

      const xhr = new XMLHttpRequest();
      xhr.withCredentials = true;

      xhr.addEventListener("readystatechange", function() {
        if (this.readyState === 4 && this.status === 401) {
          self.$router.push("/");
        }
      });

      xhr.open("GET", `${API_HOST}/auth/is_auth`);
      xhr.send();
    },
    getClasses() {
      return new Promise(resolve => {
        const xhr = new XMLHttpRequest();
        xhr.withCredentials = true;

        xhr.addEventListener("readystatechange", () => {
          if (xhr.readyState === 4) {
            if (xhr.status === 200) {
              this.classes = JSON.parse(xhr.responseText);
              this.$store.commit('setClasses', this.classes)
              resolve()
            }
          }
        });

        xhr.open("GET", `${API_HOST}/class_schedule/get_classes`);
        xhr.send();
      })
    },
    getMembers(group) {
      const xhr = new XMLHttpRequest();
      xhr.withCredentials = true;

      xhr.addEventListener("readystatechange", () => {
        if (xhr.readyState === 4) {
          if (xhr.status === 200) {
            this.members = JSON.parse(xhr.responseText);
          }
        }
      });

      xhr.open("GET", `${API_HOST}/group/members/${group}`);
      xhr.send();
    },
    getAttendance(classId) {
      const xhr = new XMLHttpRequest();
      xhr.withCredentials = true;

      xhr.addEventListener("readystatechange", () => {
        if (xhr.readyState === 4) {
          if (xhr.status === 200) {
            let attendances = JSON.parse(xhr.responseText).attendances;
            attendances.forEach(attendance => {
              this.attendance[attendance.userUid] = attendance.value === 1; // Преобразование чисел 0/1 в true/false
            })
          }
        }
      });

      xhr.open("GET", `${API_HOST}/attendance/get/${classId}`);
      xhr.send();
    },
    saveAttendance() {
      const classId = this.classes.find(value => Number(value.id) === Number(this.$route.params.classId)).id;
      const attendances = Object.entries(this.attendance).map(([userUid, value]) => ({
        class_id: classId,
        user_uid: userUid,
        value: value ? 1 : 0 // Преобразование true/false в числа 0/1
      }));

      const xhr = new XMLHttpRequest();
      xhr.withCredentials = true;

      xhr.addEventListener("readystatechange", () => {
        if (xhr.readyState === 4) {
          if (xhr.status === 200) {
            console.log('Attendance saved successfully');
          }
        }
      });

      xhr.open("POST", `${API_HOST}/attendance/set`);
      xhr.setRequestHeader("Content-Type", "application/json");
      xhr.send(JSON.stringify(attendances));
    },
    goToMainPage() {
      this.$router.push('/main'); // Переход на главную страницу (MainPage)
    }
  }
};
</script>

<style scoped>
</style>