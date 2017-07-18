<template>
  <div class="FindMatch">
    <div class="PageTop">
        <h1>This is the FindMatch Page</h1>
    </div>
    <div class="progressbar">
      <b-progress v-model="counter" :precision="1" show-progress animated></b-progress>
    </div>
    <div class="profilesection">
      <br></br>
      <h5 class="foundmatch">
        Found A Match!
      </h5>
      <br></br>
      <br></br>
      <div>
        <b-card class="profile hvr-grow">
          <h5 class="lunchbuddytitle">Lunch Buddy</h5>
          <img class="profileimg" src="./../assets/profile_picture.png">
        </b-card>
        <b-card class="restaurant-profile hvr-grow">
          <h6 class="restauranttitle">Restaurant</h6>
          <img class="restaurantprofileimg" src="./../assets/restaurant_profile.png">
        </b-card>
      </div>
      <br></br>
      <br></br>
      <br></br>
    </div>
    <div class="footer">
      Footer
    </div>
  </div>
</template>

<script>
export default {
  name: 'findmatch',
  data () {
    return {
      counter: 75
    }
  },
  methods: {
    clicked () {
      this.counter = Math.random() * 100
      console.log('Change progress to ' +
        Math.round(this.counter * 100) / 100)
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
@import './../assets/hover-min.css';
ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}
.PageTop {
  text-align: center;
  height: 400px;
  background-color: orange;
}
.foundmatch {
  text-align: center;
  font-size: 72px;
  color: red;
}
.progressbar {
}
.profilesection {
  background-image: url(./../assets/SecondaryBackground.jpg);
}
.profile {
  margin-left: 200px;
  width: 500px;
  background-image: url(./../assets/ProfileBackground.jpg);
  height: 600px;
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
}
.restaurant-profile {
  margin-left: 150px;
  width: 500px;
  background-image: url(./../assets/ProfileBackground.jpg);
  background-size: 700px 600px;
  height: 600px;
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
}
.restaurantprofileimg {

}
.profileimg {
  width: 250px;
  text-align: center;
  margin-left: 100px;
}
.lunchbuddytitle {
  font-size: 48px;
  text-align: center;
}
.restauranttitle {
  font-size: 48px;
  text-align: center;
}
</style>