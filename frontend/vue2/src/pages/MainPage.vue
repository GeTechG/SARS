<template>
  <div class="main-page">
    <h1>Расписание занятий</h1>
    <div class="table-container">
      <table class="class-table">
        <thead>
        <tr>
          <th>Номер пары</th>
          <th>Предмет</th>
          <th>Группа</th>
          <th>Дата</th>
        </tr>
        </thead>
        <tbody>
        <tr
            v-for="classItem in sortedClasses"
            :key="classItem.id"
            @click="handleClassClick(classItem.id)"
        >
          <td>{{ classItem.order }}</td>
          <td>{{ parseSubject(classItem.subject) }}</td>
          <td>{{ classItem.group }}</td>
          <td>{{ parseDate(classItem.date) }}</td>
        </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import { format, setDefaultOptions } from 'date-fns';
import { ru } from 'date-fns/locale';
import { API_HOST } from "@/config";

setDefaultOptions({ locale: ru })

export default {
  name: "MainPage",
  data() {
    return {
      classes: [],
    };
  },
  mounted() {
    this.getClasses();
  },
  computed: {
    sortedClasses() {
      return this.classes.slice().sort((a, b) => a.order - b.order);
    }
  },
  created() {
    this.checkAuthentication()
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
      const xhr = new XMLHttpRequest();
      xhr.withCredentials = true;

      xhr.addEventListener("readystatechange", () => {
        if (xhr.readyState === 4) {
          if (xhr.status === 200) {
            this.classes = JSON.parse(xhr.responseText);
            this.$store.commit('setClasses', this.classes)
          }
        }
      });

      xhr.open("GET", `${API_HOST}/class_schedule/get_classes`);
      xhr.send();
    },
    handleClassClick(classId) {
      this.$router.push(`/attendance/${classId}`);
    },
    parseDate(date) {
      const parsedDate = new Date(date);
      return format(parsedDate, "d MMMM yyyy");
    },
    parseSubject(subject) {
      return ['Математика', 'Рус.Яз', 'Java'][subject]
    }
  },
};
</script>

<style scoped>
.main-page {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.table-container {
  width: 80%;
  margin-top: 20px;
}

.class-table {
  width: 100%;
  border-collapse: collapse;
}

th, td {
  padding: 10px;
  border-bottom: 1px solid #ccc;
  text-align: center;
}

th {
  background-color: #6292ec;
  color: white;
}

tr:hover {
  background-color: #f8f9fb;
  cursor: pointer;
}


th:first-child {
  background-color: #6292ec;
  color: white;
}

td:first-child {
  font-weight: bold;
}

</style>