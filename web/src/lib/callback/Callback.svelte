<script>
  import { onMount } from "svelte";
  import { FRONT_URL } from "../api";
  import Text from "../text/Text.svelte";

  onMount(() => {
    const urlParams = new URLSearchParams(window.location.search);
    const token = urlParams.get("callback");
    const payload = { callback: token };
    window.opener.postMessage(payload, FRONT_URL);
  });
</script>

<Text key="callback" />
