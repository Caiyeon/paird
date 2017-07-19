<<template>
    <div class="result">
        <h1> Sign Up Successful! </h1>
    </div>
</template>


<script>

</script>


<style scoped>

h1, h2 {
  font-weight: normal;
}

.result {
  display: inline-block;
  position: relative;

}
.result::before {
  content: "";
  display: block;
  position:absolute;
  z-index: -1;
  width: 100%;
  height: 100%;
  background: url("./../assets/signup-5.jpeg");
  background-size: 100% 100%;
  background-repeat: no-repeat;
  opacity: 0.8;
}

</style>