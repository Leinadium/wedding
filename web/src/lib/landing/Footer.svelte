<script lang="ts">
  import github from "../../assets/github.svg";
  import Text from "../text/Text.svelte";
</script>

<footer>
  <span class="formal">
    <Text key="footer-madeby" />
  </span>
  <a href="https://github.com/Leinadium/wedding">
    <img src={github} alt="Github" />
  </a>
</footer>

<style>
  footer {
    display: flex;
    flex-flow: column wrap;
    justify-content: center;
    gap: 0.5em;
    align-items: center;
  }
  span,
  a,
  img {
    opacity: 80%;
    height: 100%;
    padding: 0;
    margin: 0;
  }
  a,
  img {
    height: 20px;
  }
</style>
