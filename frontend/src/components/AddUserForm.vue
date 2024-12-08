<template>
  <div>
    <h2>Add New Player</h2>
    
    <!-- Add User Form -->
    <form @submit.prevent="submitForm">
      <div>
        <label for="username">Username:</label>
        <input 
          type="text" 
          id="username" 
          v-model="newUser.username" 
          required
          placeholder="Enter Username" 
        />
      </div>
      
      <div>
        <label for="email">Email:</label>
        <input 
          type="email" 
          id="email" 
          v-model="newUser.email" 
          required
          placeholder="Enter Email" 
        />
      </div>
      
      <div>
        <button type="submit">Add Player</button>
      </div>
    </form>

    <!-- Success/Error Messages -->
    <div v-if="statusMessage" :class="statusClass">
      {{ statusMessage }}
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      newUser: {
        username: '',
        email: ''
      },
      statusMessage: '',  // To display success/error messages
      statusClass: ''  // To set the message color (success or error)
    };
  },
  methods: {
    // Function to submit the form
    async submitForm() {
      try {
        // Validate form data
        if (!this.newUser.username || !this.newUser.email) {
          this.statusMessage = 'Please fill in both fields!';
          this.statusClass = 'error';
          return;
        }

        // Make the POST request to add the user
        const response = await axios.post('http://localhost:3000/api/add-user', this.newUser);
        
        // Handle success
        if (response.data.success) {
          this.statusMessage = 'Player added successfully!';
          this.statusClass = 'success';
          this.resetForm();  // Clear form fields after success
        } else {
          this.statusMessage = 'Error: ' + response.data.message;
          this.statusClass = 'error';
        }
      } catch (error) {
        console.error('Error submitting form:', error);
        this.statusMessage = 'Error: Unable to add player.';
        this.statusClass = 'error';
      }
    },
    
    // Function to reset form after success
    resetForm() {
      this.newUser.username = '';
      this.newUser.email = '';
    }
  }
};
</script>

<style scoped>
/* Styling for the Add User Form */
form {
  display: flex;
  flex-direction: column;
  width: 300px;
  margin: 0 auto;
}

label {
  margin-bottom: 5px;
}

input {
  padding: 8px;
  margin-bottom: 15px;
  border-radius: 5px;
  border: 1px solid #ddd;
}

button {
  padding: 10px;
  background-color: #28a745;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

button:hover {
  background-color: #218838;
}

/* Success and error message styles */
.success {
  color: green;
}

.error {
  color: red;
}
</style>
