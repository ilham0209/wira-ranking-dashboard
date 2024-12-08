import { createRouter, createWebHistory } from 'vue-router';
import HomePage from './components/HomePage.vue';
import RankingPage from './components/RankingTable.vue';

const routes = [
  { path: '/', component: HomePage },    
  { path: '/rankings', component: RankingPage }, 
];

const router = createRouter({
  history: createWebHistory(),
  routes, 
});

export default router;
