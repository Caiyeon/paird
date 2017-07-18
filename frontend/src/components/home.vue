<template>
   <div id="Home" class="Home">
      <!-- <h1>This is the homepage</h1> -->
      <div class="separator"></div>
      <div class="row">
        <div class="Homepage row">
          <div class="col-md-3"></div>

          <div class="col-md-6">
            <div class="row Logo-title">
              <div class="col-md-4">
                <img src="../assets/PearLogo.png">
              </div>
              <div class="col-md-8">
                <h1>Pair'd</h1>
                <p>Pair'd is an app that will revolutionize the way you spend your lunch hours!</p>
                <p>First, create a profile describing your interests and lunch preferences.  The
                  app will then pair you with colleagues that share the same tastes!  Pair'd 
                  also recommends restaurants that both of you would like!  Pair'd makes
                  arranging lunch outings easy!</p>
                </p>
              </div>
            </div>
          </div>
          
          <div class="col-md-3"></div>
        </div>
      </div>
      <div class="separator"></div>
  </div>
</template>

<script>
export default {
  name: 'home'
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.Logo-title {
  margin-top: 50px;
  padding: 20px;
  background-color: #A66FA6;
  border-radius: 8px;
  opacity: 0.9;
  text-align: center
}
.separator {
  background-color:#6F256F;
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
  color: white
}
</style>
