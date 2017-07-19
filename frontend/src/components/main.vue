<template>
  <div id="app">
    <div class="Top_Panel">
      <br></br>
      <br></br>
      <div>
        <h1 class="PageTitle">Pair'd</h1>
      </div>
      <br></br>
      <div>
        <h2 class="PageSubtitle">Your Meal, Your Way</h2>
      </div>
      <br></br>
      <b-button class="GetStartedButton hvr-grow" size='lg' variant="primary" href="">
          Get Started
      </b-button>
    </div>
    <router-view></router-view>
  </div>
</template>

<script>
import VuePassword from 'vue-password'

export default {
  name: 'app',
  components: {
    VuePassword
  },
  data () {
    return {
      user: {
        email: '',
        password: ''
      },
      score: 0
    }
  },
  methods: {
  }
}
</script>

<style>
#app {
  font-family: 'noto',serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
} 
@import './../assets/hover-min.css';
.Tabs {
  font-family: Helvetica, Arial, Sans-Serif;
  font-weight: bold;
  color: black;
}
.Top_Panel {
  background-image: url("./../assets/MainBackground.jpg");
  background-size: 100% 760px;
  width: 100%;
  height: 760px;
}
.PageTitle {
font-family: 'droid';
font-style: italic;
color: white;
margin-left: 250px;
font-size: 72px;
}
.PageSubtitle {
font-family: 'Arial';
font-style: italic;
color: white;
margin-left: 200px;
font-size: 32px;
}
.GetStartedButton {
  margin-left:260px;
  font-size: 24px;
  color: white;
}
</style>
