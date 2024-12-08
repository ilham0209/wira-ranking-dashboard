<template>
  <div id="app">
    <audio autoplay loop>
      <source src="@/assets/Chase Scene.mp3" type="audio/mp3" />
      Your browser does not support the audio element.
    </audio>
    
    <MainLayout />
  </div>
</template>

<script>

import MainLayout from './components/MainLayout.vue';

export default {
  name: 'App',
  components: {
    MainLayout
  }
};
</script>

<style>
/* Gaya global untuk aplikasi */
#app {
  font-family: Arial, sans-serif;
  text-align: center;
  background-image: url('@/assets/images/image1.jpg'); 
  background-size: cover; 
  background-position: center; 
  background-attachment: fixed; 
  color: white;
  min-height: 100vh; 
  margin: 0;
  padding: 0;
}

audio {
  display: none;
}
</style>
