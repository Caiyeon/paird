<template>
  <div class="FindMatch">
    <div class="row">
        <div class="Homepage row">
          <div class="col-md-3"></div>

          <div class="col-md-6">
            <div class="row Logo-title">
              <div class="col-md-4">
                <img src="../assets/PearLogo.png">
              </div>
              <div class="col-md-8">
                <h1>Finding a Mentor/Mentee</h1>
                <p>Follow the below steps and find a restaurant and mentor/mentee that meet your preferences!</p>
                
                </br>
              
                <div class="row">
                  <div class="col-md-2">
                    <div id="circle">1</div>
                  </div>
                  <div class="col-md-8">
                    <p>If you are a new member, please create an account with us first!</p>
                    <b-button size='' variant="warning" href="" class="button-style">
                      Sign-Up
                    </b-button>
                  </div>
                </div>

                <hr>

                <div class="row">
                  <div class="col-md-2">
                    <div id="circle">2</div>
                  </div>
                  <div class="col-md-8">
                    <p>Make sure that your profile is set up correctly so we know your preferences!</p>
                    <b-button size='' variant="warning" href="" class="button-style">
                      View Profile
                    </b-button>
                  </div>
                </div>

                <hr>

                <div class="row">
                  <div class="col-md-2">
                    <div id="circle">3</div>
                  </div>
                  <div class="col-md-8">
                    <p>Your all set up! Time to meet your partner!</p>
                    <b-button size='lg' variant="success" href="" class="button-style">
                      Pair!
                    </b-button>
                  </div>
                </div>

                </p>
              </div>
            </div>
          </div>

          <div class="col-md-3">

            
          </div>
        </div>
      </div>
    <div class="progressbar">
      <b-progress v-model="counter" :precision="1" show-progress animated></b-progress>
    </div>
    <div class="profilesection">
      <br></br>
      <br></br>
      <div>
        <b-card class="profile hvr-grow">
          <h5 class="lunchbuddytitle">Mentor</h5>
          <img class="profileimg" src="./../assets/profile_picture.png">
          <br></br>
          <h7><b-badge pill variant="success">Name</b-badge>Justin Sidhu</h7>
          <br></br>
          <h7><b-badge pill variant="success">Location</b-badge>Vancouver</h7>
          <br></br>
          <h7><b-badge pill variant="success">Diet</b-badge>Vegetarian</h7>
        </b-card>
        <b-card class="restaurant-profile hvr-grow">
          <h6 class="restauranttitle">Restaurant</h6>
          <img class="restaurantprofileimg" src="./../assets/restaurant_profile.png">
          <br></br>
          <h7><b-badge pill variant="success">Name</b-badge>Cactus Club</h7>
          <br></br>
          <h7><b-badge pill variant="success">Location</b-badge>Vancouver</h7>
          <br></br>
          <h7><b-badge pill variant="success">Menu</b-badge>
            <b-button size='lg' variant="link" href="https://www.cactusclubcafe.com/menus/" class="button-style">
              See Menu!
            </b-button>
          </h7>
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
@import './../assets/hover-min.css';
h7 {
  font-size: 30px;
  margin-left: 180px;
}
ul {
  list-style-type: none;
  padding: 0;
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
  background-size: 700px 600px;
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
  width: 250px;
  text-align: center;
  margin-left: 100px;
}
.profileimg {
  width: 220px;
  text-align: center;
  margin-left: 120px;
}
.lunchbuddytitle {
  font-size: 48px;
  text-align: center;
}
.restauranttitle {
  font-size: 48px;
  text-align: center;
}
.Logo-title {
  margin-top: 50px;
  padding: 20px;
  border-radius: 8px;
  text-align: center;
  background-color: rgba(0, 0,0, 0.6);
}
.separator {
  background-color:#FFF1D6;
  height:125px;
  width:100%;
  margin: 0 auto;
}
.footer {
  background-color:#D3D3D3;
  height:50px;
  width:100%;
  margin: 0 auto;
}
.Homepage {
  height: 760px;
  width: 100%;
  background-image: url("../assets/FindMatchBackground.jpeg");
  background-size: 100% 760px;
}
.Homepage::after {
  opacity: 0.7;
}
#jumboTronStyle {
  color: #42b983;
  background-color: red;
}
h1, h2 {
  font-weight: bold;
  font-family: Arial, Helvetica, sans-serif;
  font-style: italic;
  text-align: center;
  color: white
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}

p {
  font-family: Arial, Helvetica, sans-serif;
  font-weight: bold;
  font-size: 15pt;
  color: white;
  text-align: left;
}

hr {
  border: 1px solid #ccc;
}
#circle {
  width: 50px;
  height: 50px;
  -webkit-border-radius: 25px;
  -moz-border-radius: 25px;
  border-radius: 25px;
  background: #FF970F;
  text-align: center;
  vertical-align: middle;
  line-height: 50px;
  font-weight: bold;
  font-family: Arial, Helvetica, sans-serif;
  font-size: 20;
  color: white;
}
</style>