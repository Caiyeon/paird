<template>
  <div class="signup">
    <div class="title">
      <h8>Welcome to Pair'd</h8>
    </div>

    <br> <br>

    <div class="slack_account">
      <b-form-fieldset 
      description="Link it to your Slack account."
      label="Enter your Slack account"
      :feedback="feedback_slack" 
      :state="state_slack"
      :label-cols="3"
      >
        <b-form-input v-model="slack" :state="state"></b-form-input>
      </b-form-fieldset>  
    </div> 

    <div class="profile_name">
      <b-form-fieldset 
      description="Choose your own nickname."
      label="Enter your profile name"
      :feedback="feedback_name" 
      :state="state_name"
      :label-cols="3"
      >
        <b-form-input v-model="name" :state="state"></b-form-input>
      </b-form-fieldset>  
    </div> 

    <div class="password">
      <b-form-fieldset 
      description="Password has to be at least eight characters."
      label="Enter your password"
      :feedback="feedback_password" 
      :state="state_password"
      :label-cols="3"
      >
        <b-form-input v-model="password" :state="state"></b-form-input>
      </b-form-fieldset>  
    </div> 

    <div class="password_confirm">
      <b-form-fieldset 
      label="Confirm your password"
      :feedback="feedback_confirm" 
      :state="state_confirm"
      :label-cols="3"
      >
        <b-form-input v-model="confirm" :state="state"></b-form-input>
      </b-form-fieldset>  
    </div> 

    <div class="signup_button hvr-grow">
      <b-button size="" variant="outline-success" href="">
        Sign Up
      </b-button>
    </div> 
    
    <br> <br> <br> <br> <br> <br> <br> <br> <br> <br> <br> <br> <br> <br> <br> <br> <br> <br> <br>

  </div>
</template>


<script>
/* var Form1 = document.getElementById('Form1') */

export default {
  computed: {
    feedback_slack () {
      return this.slack.length ? '' : 'Invalid'
    },
    feedback_name () {
      return this.name.length ? '' : 'Invalid'
    },
    feedback_password () {
      return this.password.length ? '' : 'Invalid'
    },
    feedback_confirm () {
      return this.confirm.length ? '' : 'Invalid'
    },
    state_slack () {
      return this.slack.length ? 'success' : 'warning'
    },
    state_name () {
      return this.name.length ? 'success' : 'warning'
    },
    state_password () {
      return this.password.length ? 'success' : 'warning'
    },
    state_confirm () {
      return this.confirm.length ? 'success' : 'warning'
    }
  },
  data () {
    return {
      slack: '',
      name: '',
      password: '',
      confirm: ''
    }
  }

}
</script>


<!-- Add "scoped" attribute to limit CSS to this component only -->

<style scoped>
@import './../assets/hover-min.css';

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

.signup {
  display: inline-block;
  position: relative;

}
.signup::before {
  content: "";
  display: block;
  position:absolute;
  z-index: -1;
  width: 100%;
  height: 100%;
  background: url("./../assets/signup-5.jpeg");
  background-size: 100% 100%;
  background-repeat: no-repeat;
  opacity: 0.6;
}

.title {
  font-family: 'Trebuchet MS', 'Lucida Sans Unicode', 'Lucida Grande', 'Lucida Sans', Arial, sans-serif;
  font-style: oblique;
  font-size: 60px;
  text-align: center;
  color: black;
  margin-top: 100px;
}

.slack_account {
  font-family: 'Trebuchet MS', 'Lucida Sans Unicode', 'Lucida Grande', 'Lucida Sans', Arial, sans-serif;
  font-size: 25px;
  position: relatives;
  margin-left: 600px;
  margin-right: 600px;
}

.profile_name {
  font-family: 'Trebuchet MS', 'Lucida Sans Unicode', 'Lucida Grande', 'Lucida Sans', Arial, sans-serif;
  font-size: 25px;
  position: relatives;
  margin-left: 600px;
  margin-right: 600px;
}

.password {
  font-family: 'Trebuchet MS', 'Lucida Sans Unicode', 'Lucida Grande', 'Lucida Sans', Arial, sans-serif;
  font-size: 25px;
  position: relatives;
  margin-left: 600px;
  margin-right: 600px;
}

.password_confirm {
  font-family: 'Trebuchet MS', 'Lucida Sans Unicode', 'Lucida Grande', 'Lucida Sans', Arial, sans-serif;
  font-size: 25px;
  position: relatives;
  margin-left: 600px;
  margin-right: 600px;
}

.signup_button {
  position: relatives;
  margin-left: 800px;
  margin-right: 800px;
}

</style>