<script>
  import Text from "../text/Text.svelte";
</script>

<header>
  <a href="/" class="date formal">
    <Text key="header-names" />
  </a>
</header>

<style>
  header {
    z-index: 1;
    text-align: center;
    margin-bottom: 2rem;
    padding-top: 2rem;

    display: flex;
    flex-flow: column nowrap;
    justify-content: start;
    align-items: center;
  }

  a {
    letter-spacing: 3px;
    text-decoration: none;
    font-size: 1.5rem;
    color: #fae7b6;
  }
</style>
