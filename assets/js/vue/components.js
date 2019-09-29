import PreviewRecipeBtn from './components/recipes/edit/PreviewRecipeBtn'
import PublishRecipeBtn from './components/recipes/edit/PublishRecipeBtn'
import DeleteRecipeBtn from './components/recipes/edit/DeleteRecipeBtn'
import CategoriesField from './components/recipes/edit/CategoriesField'
import SwitchLangModal from './components/recipes/edit/SwitchLangModal'
import DraftRecipeBtn from './components/recipes/edit/DraftRecipeBtn'
import SaveRecipeBtn from './components/recipes/edit/SaveRecipeBtn'
import LanguageSwitcher from './components/navbar/LanguageSwitcher'
import RandomRecipes from './components/recipes/RandomRecipes'
import ThemeSwitcher from './components/navbar/ThemeSwitcher'
import SortButtons from './components/recipes/SortButtons'
import LogoutBtn from './components/navbar/LogoutBtn'
import ListItem from './components/recipes/ListItem'
import AlertMessage from './components/AlertMessage'
import BtnFavs from './components/recipes/BtnFavs'
import BtnLike from './components/recipes/BtnLike'
import Recipes from './components/recipes/Recipes'
import Visibility from './components/Visibility'
import Tabs from './components/Tabs'
import Tab from './components/Tab'

Vue.component('preview-recipe-btn', PreviewRecipeBtn)
Vue.component('publish-recipe-btn', PublishRecipeBtn)
Vue.component('language-switcher', LanguageSwitcher)
Vue.component('delete-recipe-btn', DeleteRecipeBtn)
Vue.component('switch-lang-modal', SwitchLangModal)
Vue.component('categories-field', CategoriesField)
Vue.component('draft-recipe-btn', DraftRecipeBtn)
Vue.component('save-recipe-btn', SaveRecipeBtn)
Vue.component('theme-switcher', ThemeSwitcher)
Vue.component('random-recipes', RandomRecipes)
Vue.component('alert-message', AlertMessage)
Vue.component('sort-buttons', SortButtons)
Vue.component('visibility', Visibility)
Vue.component('logout-btn', LogoutBtn)
Vue.component('list-item', ListItem)
Vue.component('btn-favs', BtnFavs)
Vue.component('btn-like', BtnLike)
Vue.component('recipes', Recipes)
Vue.component('tabs', Tabs)
Vue.component('tab', Tab)