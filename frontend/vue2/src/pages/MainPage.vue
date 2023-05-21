<template>
  <div class="main-page">
    <h1>Расписание занятий</h1>
    <div class="class-list">
      <div
          v-for="classItem in classes"
          :key="classItem.date + classItem.order"
          class="class-card"
          @click="handleClassClick(classItem)"
      >
        <div class="class-card-content">
          <h3>{{ classItem.subject }}</h3>
          <p>Группа: {{ classItem.group }}</p>
          <p>Дата и время: {{ parseDate(classItem.date) }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {format, setDefaultOptions} from 'date-fns';
import {ru} from 'date-fns/locale';

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
  methods: {
    getClasses() {
      const xhr = new XMLHttpRequest();
      xhr.withCredentials = true;

      xhr.addEventListener("readystatechange", () => {
        if (xhr.readyState === 4) {
          if (xhr.status === 200) {
            this.classes = JSON.parse(xhr.responseText);
          }
        }
      });

      xhr.open("GET", "http://localhost:8001/class_schedule/get_classes");
      xhr.send();
    },
    handleClassClick(classItem) {
      // Обработка клика на карточку занятия
      console.log("Clicked on class:", classItem);
    },
    parseDate(date) {
      const parsedDate = new Date(date);
      return format(parsedDate, "d MMMM yyyy HH:mm");
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

.class-list {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 20px;
}

.class-card {
  border: 1px solid #ccc;
  border-radius: 10px;
  padding: 20px;
  cursor: pointer;
}

.class-card-content {
  text-align: center;
}
</style>