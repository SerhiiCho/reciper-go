<template>
    <div class="row py-2">
        <h3 class="col s12 center mb-4">
            {{ _recipesCategoriesTitle }} {{ fields }} / 4
        </h3>

        <div class="col s12 center mb-5" v-if="!recipeIsReady">
            <a class="add-remove-field no-select ml-2 green-text"
                :class="classAddBtn"
                @click="addField"
            >
                {{ _formsAdd + ' +' }}
            </a>
            <a class="red-text no-select add-remove-field ml-2"
                :class="classDelBtn"
                @click="deleteField"
            >
                {{ _formsDeleting + ' -' }}
            </a>
        </div>

        <div v-for="(field, i) in fields" :key="field" class="col s12 m6 mb-3">
            <div class="form-group simple-group">
                <label :for="'category_id' + field">
                    {{ _recipesCategory }} {{ field }}
                </label>

                <select name="categories[]" v-if="categories.length > 0" :disabled="recipeIsReady">
                    <option v-if="recipeCategories[i]"
                        :value="recipeCategories[i]['id']"
                        selected
                    >
                        {{ categories[recipeCategories[i]['id'] - 1]['name'] }}
                    </option>
                    <option v-for="category in categories" :key="category.id" :value="category.id">
                        {{ category.name }}
                    </option>
                </select>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            fields: 1,
            visibleAddBtn: true,
            visibleDelBtn: false,
        }
    },

    props: {
        recipeCategories: { required: false, },
        categories: { required: true, },
        _recipesCategoriesTitle: { required: true, },
        _formsAdd: { required: true, },
        _formsDeleting: { required: true, },
        _recipesCategory: { required: true, },
        recipeIsReady: { required: true, }
    },

    created() {
        this.getFieldsFromProps()
    },

    methods: {
        getFieldsFromProps() {
            if (this.recipeCategories) {
                if (this.recipeCategories.length == 0) {
                    this.fields = 1
                } else {
                    this.fields = this.recipeCategories.length
                }
                this.stableButtons()
            }
        },

        addField() {
            if (this.fields <= 3 && this.fields > 0) {
                this.fields++
            }
            this.stableButtons()
        },

        deleteField() {
            if (this.fields > 1) {
                this.fields--
            }
            this.stableButtons()
        },

        stableButtons() {
            if (this.fields > 1) {
                this.visibleDelBtn = true
            }

            if (this.fields === 4) {
                this.visibleAddBtn = false
                this.visibleDelBtn = true
            }

            if (this.fields < 4) {
                this.visibleAddBtn = true
            }

            if (this.fields === 1) {
                this.visibleDelBtn = false
            }
        },
    },
    computed: {
        classDelBtn() {
            return {
                disable: !this.visibleDelBtn,
            }
        },
        classAddBtn() {
            return {
                disable: !this.visibleAddBtn,
            }
        },
    },
}
</script>