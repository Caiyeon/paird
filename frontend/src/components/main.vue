<template>
  <div id="app">
    <div class="Top_Panel">
      <br></br>
      <div class="row PageTitle">
        <div class="col-md-2">
        </div>
        <div class="col-md-2">
            <img src="./../assets/MainPage_Pearlogo.png"/>
        </div>
        <div class="col-md-1">
            <br>
            <h8>Pair'd</h8>
        </div>
        <div class="col-md-1">
            <img src="./../assets/MainPage_Pearlogo.png"/>
        </div>
      </div>
      <br></br>
      <div>
        <h2 class="PageSubtitle">Find Your Mentor Today</h2>
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
  position:absolute;
  background-image: url("./../assets/MainBackground.jpg");
  width: 100%;
  height: 100%;
  background-size: 100% 100%;
  opacity: 0.8;
}
.PageTitle {
position: relatives;
font-family: 'droid';
font-style: italic;
color: white;
margin-left: 330px;
font-size: 72px;
}
.PageSubtitle {
position: relatives;
font-family: 'Arial';
font-style: italic;
color: white;
margin-top: 100px;
margin-left: 700px;
font-size: 32px;
}
.GetStartedButton {
  position: relatives;
  margin-left: 770px;
  font-size: 24px;
  color: white;
}
</style>
