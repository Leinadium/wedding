<script lang="ts">
  import github from "../../assets/github.svg";
  import Text from "../text/Text.svelte";

  import Countdown from "../Countdown.svelte";
</script>

<footer>
  <Countdown />
  <a class="formal" href="https://github.com/Leinadium/wedding">
    <img src={github} alt="Github" />
    <Text key="footer-madeby" />
  </a>
</footer>

<style>
  footer {
    width: 100vw;
    height: 14rem;

    display: flex;
    flex-flow: column wrap;
    justify-content: flex-end;
    gap: 2rem;
    align-items: center;

    color: #fae7b6;
    background-image: url("../../assets/venue/background.png");
  }
  a {
    margin-top: 2rem;
    margin-bottom: 2rem;
    color: #fae7b6;
    text-decoration: none;
    font-size: 1rem;

    display: inline-flex;
    align-items: center;
    justify-content: center;
  }
  img {
    width: 1rem;
    margin-right: 0.5rem;
  }
</style>
