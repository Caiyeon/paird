<template>
  <div>
    <div class="separator"></div>
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
            Sumbit
          </b-button>
        </div> 

      </div>

      <div class="col-md-4">
        </br>
        </br>
        </br>
        </br>
        </br>
        </br>
        </br>
        </br>
        </br>
        </br>
        </br>
        </br>
        </br>
        </br>
        </br>
        </br>
        </br>
        </br>
        </br>
        </br>
        </br>
        </br>
        </br>
        </br>
        <div class="row">
          <div class="col-md-6">
            <h4>Pair'd Team</h4>
            <p><i class="fa fa-building-o" style="color:white"></i>
              &nbsp; SAP Vancouver
              </br>&nbsp; &nbsp; &nbsp; 910 Mainland St
              </br>&nbsp; &nbsp; &nbsp; Vancouver, BC V6B 1A9
            </p>
          </div>
          <div class="col-md-6">
            <p>
              </br><i class="fa fa-phone" style="color:white"></i> &nbsp; (604) 647-8888
              </br><i class="fa fa-envelope" style="color:white"></i> &nbsp; pair'd@sap.com
            </p>
          </div>
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

h1, h2, h4{
  font-weight: initial;
  font-family: Arial, Helvetica, sans-serif;
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

p {
  font-family: Arial, Helvetica, sans-serif;
  font-size: 14pt;
  color: white;
}
</style>