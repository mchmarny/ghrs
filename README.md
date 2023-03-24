# ghstore

GitHub Actions demo for in repo, persistent data store using [SQLite](https://sqlite.org/index.html).

```yaml

- name: Checkout
    uses: actions/checkout@8f4b7f84864484a7bf31766abe9204da3cbe65b3  # v3.5.0

- name: Get
    id: get
    uses: mchmarny/ghstore@main
    with:
    state: ${{ inputs.data_file }}  # will be created if doesn't exist
    key: ${{ inputs.data_key }}
    operation: get

- name: Print Get Output
    run: |-
    set -euo pipefail
    echo "Current value: ${{ steps.get.outputs.value }}"

- name: Add
    uses: mchmarny/ghstore@main
    with:
    state: ${{ inputs.data_file }}
    key: ${{ inputs.data_key }}
    operation: add
    value: '10'

- name: Sub
    id: sub
    uses: mchmarny/ghstore@main
    with:
    state: ${{ inputs.data_file }}
    key: ${{ inputs.data_key }}
    operation: sub
    value: '3'

- name: Print Final Output
    run: |-
    set -euo pipefail
    echo "Current value: ${{ steps.sub.outputs.value }}"

- name: Save
    run: |
    git config --global user.name ${{ github.actor }}
    git config --global user.email ${{ github.actor }}@users.noreply.github.com
    git add ${{ inputs.data_file }}
    git commit -am "save ${{ inputs.data_file }}"
    git push origin main
```



## disclaimer

This is my personal project and it does not represent my employer. While I do my best to ensure that everything works, I take no responsibility for issues caused by this code.