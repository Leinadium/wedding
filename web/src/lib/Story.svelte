<script lang="ts">
  import { fade } from "svelte/transition";
  import p0 from "../assets/story/p0.png";
  import p1 from "../assets/story/p1.png";
  import p2 from "../assets/story/p2.png";
  import p3 from "../assets/story/p3.png";
  import p4 from "../assets/story/p4.png";
  import p5 from "../assets/story/p5.png";

  import { Page } from "./common";
  import Image from "./Image.svelte";
  import Text from "./text/Text.svelte";
  import { onDestroy, onMount } from "svelte";
  let {
    setState,
  }: {
    setState: (page: Page) => void;
  } = $props();

  function back() {
    setState(Page.Landing);
  }

  function scrollTop() {
    window.scrollTo({ top: 0, behavior: "smooth" });
  }

  onMount(scrollTop);
  onDestroy(scrollTop);
</script>

<div class="story">
  <div id="sep" class="photo">
    <Image src={p0} alt="p0" delay={0} y={100} duration={100} />
  </div>

  <p class="formal" in:fade={{ delay: 400, duration: 200 }}>
    &emsp;&emsp;<Text key="story-text01" /><br />
    &emsp;&emsp;<Text key="story-text02" /><br />
    &emsp;&emsp;<Text key="story-text03" /><br />
    &emsp;&emsp;<Text key="story-text04" /><br />
    &emsp;&emsp;<Text key="story-text05" /><br />
  </p>

  <div id="p1" class="photo" style="transform: rotate(5deg);">
    <Image src={p1} alt="p1" delay={600} y={100} duration={1000} />
  </div>

  <p class="formal" in:fade={{ delay: 800, duration: 200 }}>
    &emsp;&emsp;<Text key="story-text06" /><br />
    &emsp;&emsp;<Text key="story-text07" /><br />
  </p>

  <div id="p2" class="photo" style="transform: rotate(-5deg);">
    <Image src={p2} alt="p2" delay={1000} y={100} duration={1000} />
  </div>

  <p class="formal" in:fade={{ delay: 1200, duration: 200 }}>
    &emsp;&emsp;<Text key="story-text08" /><br />
    &emsp;&emsp;<Text key="story-text09" /><br />
    &emsp;&emsp;<Text key="story-text10" /><br />
  </p>

  <div id="p3" class="photo" style="transform: rotate(5deg);">
    <Image src={p3} alt="p3" delay={1400} y={100} duration={1000} />
  </div>

  <p class="formal" in:fade={{ delay: 1600, duration: 200 }}>
    &emsp;&emsp;<Text key="story-text11" /><br />
    &emsp;&emsp;<Text key="story-text12" /><br />
  </p>

  <div id="p4" class="photo" style="transform: rotate(-5deg);">
    <Image src={p4} alt="p4" delay={1800} y={100} duration={1000} />
  </div>

  <p class="formal quote" in:fade={{ delay: 2000, duration: 200 }}>
    &emsp;&emsp;<Text key="story-text13" /><br />
    &emsp;&emsp;&emsp;&emsp; - <Text key="story-text14" /><br />
    &emsp;&emsp;<Text key="story-text15" /><br />
    &emsp;&emsp;&emsp;&emsp; - <Text key="story-text16" /><br />
  </p>

  <div id="p5" class="photo" style="transform: rotate(-5deg);">
    <Image src={p5} alt="p5" delay={2200} y={100} duration={1000} />
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
    display: flex;
    flex-flow: column nowrap;
    align-items: center;

    min-height: 90vh;
    width: 85vw;
    max-width: 600px;

    color: #fae7b6;
    font-size: 1.2rem;
    text-align: justify;
  }

  .photo {
    height: auto;
    width: 15vh;
  }

  #sep {
    width: 30%;
  }

  .quote {
    font-style: italic;
  }

  #go-back-wrapper {
    text-align: center;
    margin-bottom: 2rem;
    display: inline-block;
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
