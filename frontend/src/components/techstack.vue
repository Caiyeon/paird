<template>
  <div>
    <div class="ContactUs">

    <div class="row">
      <div class="col-md-6 form-class">

        <div class="heading-title">
          <h1>Contact Us</h1>
        </div>

        <b-form-fieldset label="Enter your name">
          <b-form-input v-model="name" :state="state"></b-form-input>
        </b-form-fieldset>

        <b-form-fieldset label="Enter your email">
          <b-form-input v-model="name" :state="state"></b-form-input>
        </b-form-fieldset>  

        <b-form-fieldset label="Enter your message">
          <b-form-input v-model="name" :state="state" id="message-box"></b-form-input>
        </b-form-fieldset>  

        <div class="signup_button hvr-grow">
          <b-button size="" variant="success" href="">
            Sign Up
          </b-button>
        </div> 

      </div>
    </div>

    </div>
  </div>
</template>

<script>
export default {
  name: 'contactus'
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#message-box {
  height: 200px
}
.form-class {
  font-family: Arial, Helvetica, sans-serif;
  font-size: 14pt;
  color: white;
  margin-left: 50px;
  padding: 50px 90px;
  
}
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