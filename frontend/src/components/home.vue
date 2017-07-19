<template>
   <div id="Home" class="Home">
      <!-- <h1>This is the homepage</h1> -->
      <div class="row">
        <div class="Homepage row">
          <div class="col-md-3"></div>

          <div class="col-md-6">
            <div class="row Logo-title">
              <div class="col-md-4">
                <img src="../assets/PearLogo.png">
              </div>
              <div class="col-md-8">
                <h1>Getting Started</h1>
                <p>Pair'd is an app that will revolutionize the way you spend your lunch hours!</p>
                
                </br>
              
                <div class="row">
                  <div class="col-md-2">
                    <div id="circle">1</div>
                  </div>
                  <div class="col-md-8">
                    <p>Create a profile describing your interests and lunch preferences.</p>
                    <b-button size='' variant="primary" href="" class="button-style">
                      Profile
                    </b-button>
                  </div>
                </div>

                <hr>

                <div class="row">
                  <div class="col-md-2">
                    <div id="circle">2</div>
                  </div>
                  <div class="col-md-8">
                    <p>Go to the "Find Lunch Buddy" page and click "Find Match".</p>
                    <b-button size='' variant="primary" href="" class="button-style">
                      Find Lunch Buddy
                    </b-button>
                  </div>
                </div>

                <hr>

                <div class="row">
                  <div class="col-md-2">
                    <div id="circle">3</div>
                  </div>
                  <div class="col-md-8">
                    <p>Pair'd will match you up with another colleague and will
                      recommend a restaurant based on your food preferences.</p>
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
  </div>
</template>

<script>
export default {
  name: 'home'
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.button-style {
  background-color: #008CBA;
  opacity: 1;
  width: 75%;
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
.Logo-title {
  margin-top: 50px;
  padding: 20px;
  border-radius: 8px;
  text-align: center;
  background-color: rgba(0, 0, 0, 0.6);
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
  background-image: url("../assets/HomeBackground.jpg");
  background-size: 100% 760px;
  width: 100%;
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
</style>
