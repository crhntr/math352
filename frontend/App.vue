<template>
	<div>
		<div class="page">
			<h2>STEP 1: Load Data From PubMed</h2>
			<p>This step request articles from PubMed using a pubmed query ({{query}}). The pubmed api simply searches a database of articles by date added. The goal of this project is to filter these using a machine learning algorithm based on Naive Bayes classification.
			</p>
			<button @click="fetch()" v-if="show_fetch">Fetch</button>
			<button @click="updateItems()" v-if="!show_fetch">Update Items</button>
			<div>
				<div v-for="item in items" class="article">
					<h3 @click="item.show_body = !item.show_body">{{item.data.title}}</h3>
					<p v-if="item.show_body">{{item.data.body}}</p>
					<p v-if="item.show_body">{{item.data.categories}}</p>
				</div>
			</div>
		</div>
		<div class="page classifier">
			<h2>STEP 2: Classify Articles</h2>
			<article>
				<h2>{{item.data.title}}</h2>
				<p>{{item.data.body}}</p>
			</article>
			<div class="widget">
				<div v-for="(score, cat) in item.data.categories">
					<input type="checkbox" :value="cat" :id="cat" v-model="item.categories">
					<label :for="cat">{{cat}}</label>
				</div>
				<input type="text" @keyup.enter="categories.push(newCategory)" v-model="newCategory" placeholder="Category"/>
				<br />
				<button @click="submitClasses(); next(); set_default_categories()">Next</button>
			</div>
		</div>
		<div class="page">
			<a href="http://github.com/crhntr/math352">Source code (github)</a>
		</div>
	</div>
</template>
<script>
	export default {
	  data() {
	    return {
				query: "talimogene laherparepvec  [All Fields]",
	      threshold: 0.02,
	      newCategory: '',
				categories: [],
				items: [],
				show_fetch: true,
				current_item_id: 0,
				show_score: 0,
	      item: {
	        data: {},
					id: 0,
	      },
				query: "talimogene laherparepvec [All Fields]"
	    }
	  },
	  created() {},
	  methods: {
			submitClasses() {
				if (this.item.categories.length > 0) {
					this.$http.patch("/api/item/"+this.current_item_id+"/classes", {
						"classes": this.item.categories,
					})
				}
			},
			updateItems() {
				this.$http.get("/api/items").then(response => {
					this.items = []
					for (let item of response.data.data) {
						item.body = item.body || "(no body)"
						if (Object.keys(item.categories).length === 0) {
							this.show_score = false
							item.categories['relevant'] = 0
							item.categories['irrelevant'] = 0
						}
						this.items.push({
							show_body: false,
							data: item,
							categories: []
						})
					}
					this.current_item_id = 0;
					this.item = this.items[this.current_item_id]
				}, response => {
					console.log(response)
				})
			},
			fetch() {
				this.show_fetch = false
				this.$http.post("/act/item/fetch", this.query, {params:{"days": this.weeks_back}}).then(response => {
					this.show_fetch = false
					this.updateItems()
				}, response => {
					this.show_fetch = true
				})
			},
 	    next() {
				this.current_item_id++
				if (this.current_item_id >= this.items.length) {
					this.current_item_id = 0
					alert("you have seen every article and now are starting again")
				}
				console.log("next: " + this.current_item_id)
				this.item = this.items[this.current_item_id]
	    },
			set_default_categories() {
				console.log(this.item)
				for(let cl in this.item.data.categories) {
					console.log(cl, this.item.data[cl])
				}
			},
			refreshCategory(category) {

			}
	  }
	}
</script>
