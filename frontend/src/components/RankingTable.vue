<template>
  <div>
    <h2>WIRA TOP PLAYER</h2>
    <table class="ranking-table">
      <thead>
        <tr>
          <th>Rank</th>
          <th>Username</th>
          <th>Score</th>
        </tr>
      </thead>
      <tbody>

        <tr v-for="(player, index) in players" :key="index">
          <td>{{ (currentPage - 1) * limit + index + 1 }}</td> 
          <td>{{ player.username }}</td>
          <td>{{ player.score }}</td>
        </tr>
      </tbody>
    </table>
    <div class="pagination">
      <button class="pagination-btn prev-btn" @click="goToPage(currentPage - 1)" :disabled="currentPage <= 1">Previous</button>
      <button class="pagination-btn next-btn" @click="goToPage(currentPage + 1)">Next</button>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      players: [],
      currentPage: 1,
      limit: 30,  
    };
  },
  created() {
    this.fetchPlayers();
  },
  methods: {
    async fetchPlayers() {
      try {
        const response = await axios.get(`http://localhost:8080/api/rankings?limit=${this.limit}&page=${this.currentPage}`);
        console.log("Fetched players:", response.data);  
        this.players = response.data;  
      } catch (error) {
        console.error("Error fetching players:", error);
      }
    },
    goToPage(page) {
      if (page > 0) {
        this.currentPage = page;
        this.fetchPlayers(); 
      }
    }
  }
};
</script>

<style scoped>
.ranking-table {
  width: 50%;  
  border-collapse: collapse;
  margin: 0 auto;
  margin-bottom: 20px;
  background-color: #333;
  border-radius: 10px;
}

th, td {
  padding: 12px;
  border: 1px solid #444;
  text-align: center;
  background-color: #000000;
  color: #ffffff;
  font-weight: bold;
}

h2 {
  text-align: center;
  font-size: 2rem;
  margin-bottom: 20px;
  color: #fff;
}

.pagination {
  text-align: center;
  margin-top: 30px;
}

.pagination-btn {
  background-color: #555;
  color: white;
  border: none;
  padding: 12px 25px;
  font-size: 1.2rem;
  border-radius: 5px;
  cursor: pointer;
  transition: all 0.3s ease;
  margin: 0 20px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.3);
}

.pagination-btn:hover {
  background-color: #888;
  transform: scale(1.1);
  box-shadow: 0 6px 15px rgba(0, 0, 0, 0.4);
}

.pagination-btn:disabled {
  background-color: #444;
  cursor: not-allowed;
}

.pagination-btn:focus {
  outline: none;
}

.prev-btn {
  background-color: #000000;
}

.prev-btn:hover {
  background-color: #333;
}

.next-btn {
  background-color: #e3dc56;
}

.next-btn:hover {
  background-color: #b2a100;
}

@media (max-width: 768px) {
  .ranking-table {
    font-size: 14px; 
    width: 70%;
  }

  th, td {
    padding: 10px;  
  }

  .pagination-btn {
    padding: 8px 16px; 
  }
}
</style>
