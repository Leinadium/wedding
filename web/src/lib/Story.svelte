<script lang="ts">
  import p1 from "../assets/story/p1.png";

  import { Page } from "./common";
  import Image from "./Image.svelte";
  import Text from "./text/Text.svelte";
  let {
    setState,
  }: {
    setState: (page: Page) => void;
  } = $props();

  function back() {
    setState(Page.Landing);
  }
</script>

<div class="story">
  <p id="t1" class="formal">
    A lot of text A lot of text A lot of text A lot of text A lot of text A lot
    of text A lot of text A lot of text A lot of text A lot of text A lot of
    text A lot of text A lot of text A lot of text A lot of text A lot of text A
    lot of text A lot of text A lot of text A lot of text A lot of text A lot of
    text A lot of text.
  </p>
  <p id="t2" class="formal">
    A lot of text A lot of text A lot of text A lot of text A lot of text A lot
    of text A lot of text A lot of text A lot of text A lot of text A lot of
    text A lot of text A lot of text A lot of text A lot of text A lot of text A
    lot of text A lot of text A lot of text A lot of text A lot of text A lot of
    text A lot of text A lot of text A lot of text A lot of text A lot of text A
    lot of text.
  </p>
  <p id="t3" class="formal">
    A lot of text A lot of text A lot of text A lot of text A lot of text A lot
    of text A lot of text A lot of text A lot of text A lot of text A lot of
    text A lot of text A lot of text A lot of text A lot of text A lot of text A
    lot of text A lot of text A lot of text A lot of text A lot of text A lot of
    text A lot of text A lot of text A lot of text A lot of text A lot of text A
    lot of text A lot of text A lot of text A lot of text A lot of text A lot of
    text A lot of text A lot of text A lot of text A lot of text A lot of text A
    lot of text A lot of text A lot of text A lot of text A lot of text A lot of
    text A lot of text A lot of text A lot of text A lot of text A lot of text A
    lot of text A lot of text A lot of text A lot of text A lot of text A lot of
    text A lot of text
  </p>
  <p class="formal">
    A lot of text A lot of text A lot of text A lot of text A lot of text A lot
    of text A lot of text A lot of text A lot of text A lot of text A lot of
    text A lot of text A lot of text A lot of text A lot of text A lot of text A
    lot of text A lot of text A lot of text A lot of text A lot of text A lot of
    text A lot of text A lot of text A lot of text A lot of text A lot of text A
    lot of text A lot of text A lot of text A lot of text A lot of text A lot of
    text A lot of text A lot of text A lot of text A lot of text A lot of text A
    lot of text A lot of text A lot of text A lot of text A lot of text A lot of
    text A lot of text A lot of text A lot of text A lot of text A lot of text A
    lot of text A lot of text A lot of text A lot of text A lot of text A lot of
    text A lot of text
  </p>

  <div id="p1" class="photo">
    <Image src={p1} alt="p1" delay={100} y={120} duration={1000} />
  </div>
  <div id="p2" class="photo">
    <Image src={p1} alt="p2" delay={300} y={120} duration={1230} />
  </div>

  <div id="go-back-wrapper">
    <a
      id="go-back"
      class="formal"
      href="/#"
      onclick={(e) => {
        e.preventDefault();
        back();
      }}
    >
      <Text key="story-back" />
    </a>
  </div>
</div>

<style>
  .story {
    position: relative;
    display: flex;
    flex-flow: column nowrap;

    min-height: 90vh;
    width: 100%;
    max-width: 500px;

    color: #fae7b6;
  }

  .photo {
    position: absolute;
    height: auto;
  }

  p,
  a {
    font-size: 1.1rem;
  }

  #t1 {
    width: 93%;
    align-self: flex-end;
  }

  #p1 {
    left: -30%;
    top: 1%;
    width: 32%;
    transform: rotate(-10deg);
    z-index: 1;
  }

  #t2 {
    width: 85%;
    align-self: flex-start;
  }

  #p2 {
    right: -20%;
    top: 16%;
    width: 34%;
    transform: rotate(12deg);
    z-index: 1;
  }

  #go-back-wrapper {
    text-align: center;
    margin-bottom: 2rem;
    margin-top: 2rem;
  }

  #go-back {
    text-decoration: none;
    color: #fae7b6;
    font-weight: 300;
    border: 1px solid #fae7b6;
    border-radius: 10px;
    padding: 1rem;
  }
</style>
