<template>
  <div class="main-layout">
    <audio autoplay loop muted>
      <source src="@/assets/Chase Scene.mp3" type="audio/mpeg" />
      Your browser does not support the audio element.
    </audio>

    <header>
      <h1 class="moving-text">Wira Ranking Dashboard</h1>
    </header>

    <section class="image-gallery">
      <div class="image-card" v-for="image in images" :key="image.id">
        <img :src="image.src" :alt="'Image ' + image.id" class="image" />
        <div class="image-text">{{ getImageText(image.id) }}</div> 
      </div>
    </section>


    <router-view /> 
  </div>
</template>

<script>
export default {
  data() {
    return {
      images: [
        { id: 1, src: require('@/assets/images/image4.png') },
        { id: 2, src: require('@/assets/images/image2.png') },
        { id: 3, src: require('@/assets/images/image3.png') }
      ]
    };
  },
  methods: {
    getImageText(id) {
      const textMap = {
        1: 'Xornnn Our Top 3 Player',
        2: 'Stormie Our Top 1 Player',
        3: 'Cikugais Our Top 2 Player'
      };
      return textMap[id] || '';
    }
  }
};
</script>

<style scoped>
.main-layout {
  font-family: Arial, sans-serif;
  background-image: url('@/assets/images/image1.jpg');
  background-position: center;
  background-attachment: fixed;
  color: white;
  min-height: 200vh;
}

.moving-text {
  font-size: 4rem;
  font-weight: bold;
  animation: moveText 30s linear infinite;
  text-align: center;
  padding: 1px;
}

@keyframes moveText {
  0% {
    transform: translateX(-100%);
  }
  50% {
    transform: translateX(0);
  }
  100% {
    transform: translateX(100%);
  }
}

header {
  background-color: rgba(0, 0, 0, 0.5); 
  color: white;
  padding: 10px;
  text-align: center;
}

header h1 {
  font-size: 3rem;
  font-weight: 600;
  text-shadow: 2px 2px 5px rgba(0, 0, 0, 0.7); 
}

.image-gallery {
  display: flex;
  justify-content: center;
  align-items: center; 
  gap: 40px; 
  padding: 10px;
  height: auto; 
}

.image-card {
  background-color: #222;
  border-radius: 12px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  padding: 20px;
  text-align: center;
  position: relative; 
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.image-card:hover {
  transform: translateY(-10px); 
  box-shadow: 0 6px 15px rgba(0, 0, 0, 0.3); 
}

.image {
  width: 100%;
  height: 100%;
  max-width: 300px;
  max-height: 300px; 
  object-fit: cover; 
  border-radius: 10px;
}

.image-text {
  position: absolute;
  top: 50%; 
  left: 50%; 
  transform: translate(-50%, -50%); 
  color: white;
  font-size: 2rem; 
  font-weight: bold;
  text-shadow: 2px 2px 5px rgba(0, 0, 0, 0.7); 
}

@media (max-width: 768px) {
  .moving-text {
    font-size: 2.5rem; 
  }

  .image-text {
    font-size: 1.0rem; 
  }
  
  .image {
    width: 100%;  
    max-width: 200px; 
  }

  .image-gallery {
    flex-wrap: nowrap;  
    justify-content: center; 
    gap: 20px; 
  }

  .pagination-btn {
    padding: 8px 16px;  
  }
}
</style>
