<template>
  <div id="app">
      <b-navbar toggleable type="inverse" variant="warning">
        <div>
          <router-link to="/">
            <b-button class="TopPanelButtons" size='' variant="link" href="">
              Home
            </b-button>
          </router-link>
          <router-link to="/getstarted">
            <b-button class="TopPanelButtons" size='' variant="link" href="">
              Get Started
            </b-button>
          </router-link>
          <router-link to="/profile">
            <b-button class="TopPanelButtons" size='' variant="link" href="">
              Profile
            </b-button>
          </router-link>
          <router-link to="/findmatch">
            <b-button class="TopPanelButtons" size='' variant="link" href="">
              Find Lunch Buddy
            </b-button>
          </router-link>
          <router-link to="/contactus">
            <b-button class="TopPanelButtons" size='' variant="link" href="">
              Contact Us
            </b-button>
          </router-link>
          <router-link to="/techstack">
            <b-button class="TopPanelButtons" size='' variant="link" href="">
              Tech Stack
            </b-button>
          </router-link>
          <b-button class="SignInButton hvr-grow" size='lg' variant="success" href="">
            Log In
          </b-button>
          <router-link to="/signup">
            <b-button class="SignUpButton hvr-grow" size='lg' variant="danger" href="">
              Sign Up
            </b-button>
          </router-link>
        </div>
      </b-navbar>
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
@import './assets/hover-min.css';
.Tabs {
  font-family: Helvetica, Arial, Sans-Serif;
  font-weight: bold;
  color: black;
}
.Top_Panel {
  background-image: url("./assets/MainBackground.jpg");
  background-size: 100% 760px;
  width: 100%;
  height: 760px;
}
.TopPanelButtons {
  margin-left:30px;
  font-size: 24px;
  color: white;
}
.SignInButton {
  margin-left: 100px;
  font-weight: bold;
}
.SignUpButton {
  margin-left: 20px;
  font-weight: bold;
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
