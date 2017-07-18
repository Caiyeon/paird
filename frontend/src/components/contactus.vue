<template>
  <div>
    <div class="separator"></div>
    <div class="ContactUs">

    <div class="row">
      <div class="col-md-6">
        <div class="heading-title">
        <h1>Contact Us</h1>
        </div>
      </div>
    </div>

    </div>
    <div class="separator"></div>
  </div>
</template>

<script>
export default {
  name: 'contactus'
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.heading-title {
  text-align: center;
  margin-top: 50px
}
.ContactUs {
  background-image: url("../assets/ContactUsBackground.jpeg");
  background-size: 100% 760px;
  height: 760px;
}
.separator {
  background-color:#FFF1D6;
  height:125px;
  width:100%;
  margin: 0 auto;
}

h1, h2 {
  font-weight: initial;
  font-family: Arial, Helvetica, sans-serif;
  font-size: 80px;
  font-style: italic;
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
</style>