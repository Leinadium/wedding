<script lang="ts">
  let {
    src,
    text,
    offset,
  }: {
    src: string;
    text: string;
    offset: string;
  } = $props();
</script>

<div class="header" style="--offset: {offset}">
  <img {src} alt="header" />
  <span class="cursive">{text}</span>
</div>

<style>
  div {
    position: relative;
    width: 100%;
  }

  img {
    display: block;
    width: 100%;
    min-height: 20vh;
    max-height: 50vh;
    object-fit: cover;
    z-index: 1;
  }

  span {
    font-size: max(10vw, 2rem);
    color: #f9f9f9;

    position: absolute;
    top: 50%;
    padding-top: var(--offset);
    left: 50%;
    transform: translate(-50%, -50%);
    z-index: 2;
  }
</style>
