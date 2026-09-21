<script lang="ts">
  import { load, type Payload } from "./common";

  let {
    setContent,
  }: {
    setContent: (content: Payload) => void;
  } = $props();

  let value: string = $state("");
  let error: boolean = $state(false);

  async function handle(e: Event) {
    e.preventDefault();
    if (value) {
      let res = await load(value);
      if (res) {
        setContent(res);
      } else {
        error = true;
      }
    }
  }
</script>

<section class="p-strip">
  <div class="row">
    <div class="col-4 prefix-4">
      <div class="p-card">
        <h3>Authentication</h3>
        <hr />
        <form class="p-form" onsubmit={handle}>
          <fieldset>
            <label for="pass">Password</label>
            <input
              id="pass"
              type="password"
              placeholder="Enter password..."
              bind:value
              required
            />
            {#if error}
              <p
                class="p-form-validation__message"
                id="exampleInputErrorMessage"
              >
                Could not login.
              </p>
            {/if}
            <button
              type="submit"
              class="p-button--positive u-no-margin--bottom"
            >
              Unlock
            </button>
          </fieldset>
        </form>
      </div>
    </div>
  </div>
</section>
