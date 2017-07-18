<template>
   <div id="Home" class="Home">
      <!-- <h1>This is the homepage</h1> -->
      <div class="separator"></div>
      <div class="row">
        <div class="Homepage row">
          <div class="col-md-3"></div>
          <div class="col-md-6">
            <div class="Logo-title">
              <h1>Pair'd</h1>
              <p>Pair'd is an app that will revolutionize the way you spend your lunch hours!</p>
              <p>Users create a profile describing their interests and lunch preferences and
                the app will pair 
              </p>
            </div>
          </div>
          <div class="col-md-3"></div>
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
.Logo-title {
  margin-top: 50px;
  padding: 10px;
  background-color: red;
  border-radius: 5px
}
.separator {
  background-color:#6F256F;
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
  font-weight: normal;
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
</style>
