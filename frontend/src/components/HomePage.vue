<template>
  <div>

    <div class="search-bar">
      <input
        type="text"
        v-model="searchQuery"
        @input="searchPlayer"
        placeholder="Search Player"
        class="search-input"
      />
    </div>

    <div v-if="selectedPlayer" class="popup-overlay">
      <div class="popup">
        <h2>Player Details</h2>
        <p><strong>Username:</strong> {{ selectedPlayer.username }}</p>
        <p><strong>Email:</strong> {{ selectedPlayer.email }}</p>
        <p><strong>Class ID:</strong> {{ selectedPlayer.class_id }}</p>
        <p><strong>Score:</strong> {{ selectedPlayer.reward_score }}</p>
        <button @click="closePopup">Close</button>
      </div>
    </div>

    <RankingTable :players="players" />
  </div>
</template>

<script>
import axios from 'axios';
import RankingTable from './RankingTable.vue';

export default {
  components: {
    RankingTable,
  },
  data() {
    return {
      players: [], 
      searchQuery: "", 
      selectedPlayer: null, 
    };
  },
  methods: {
    async searchPlayer() {
      if (this.searchQuery.trim() === "") {
        this.selectedPlayer = null; 
        this.loadPlayers(); 
        return;
      }

      try {
        const response = await axios.get('/api/rankings', {
          params: {
            search: this.searchQuery,
            limit: 10, 
          },
        });

        console.log('Search Results:', response.data);

        const foundPlayer = response.data.find(player => 
          player.username.toLowerCase().includes(this.searchQuery.toLowerCase())
        );

        if (foundPlayer) {
          this.selectedPlayer = foundPlayer; 
        } else {
          this.selectedPlayer = null;
        }
      } catch (error) {
        console.error("Error fetching player data:", error);
      }
    },
    closePopup() {
      this.selectedPlayer = null;
    },

    async loadPlayers() {
      try {
        const response = await axios.get('/api/rankings', { params: { limit: 10 } });
        this.players = response.data;
      } catch (error) {
        console.error("Error fetching ranking data:", error);
      }
    },
  },
  mounted() {
    this.loadPlayers(); 
  },
};
</script>

<style scoped>
.search-bar {
  display: flex;
  justify-content: center;
  margin: 50px 0; 
}

.search-input {
  padding: 10px;
  font-size: 1.2rem;
  width: 250px;
  border-radius: 5px;
  border: 1px solid #ccc;
}


.popup-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
}

.popup {
  background-color: white;
  padding: 20px;
  border-radius: 10px;
  width: 300px;
  text-align: center;
}

button {
  padding: 10px 15px;
  border-radius: 5px;
  background-color: #007bff;
  color: white;
  border: none;
  cursor: pointer;
}

button:hover {
  background-color: #0056b3;
}
</style>
