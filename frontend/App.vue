<template>
	<div>
		<div class="page">
			<h1>About Naive Bayes Classification</h1>
			<p>
				In machine learning, naive Bayes classifiers are a family of simple probabilistic classifiers based on applying Bayes' theorem with strong (naive) independence assumptions between the features.
			</p>
			<div height="100vh">
				<img src="/src/theorem.png" height="100px"/>
			</div>
		</div>
		<div class="page">
			<h1>STEP 1: Load Data From PubMed</h1>
			<p>
				This step request articles from PubMed using a pubmed query ({{query}}). The pubmed api simply searches a database of articles by date added. The goal of this project is to filter these using a machine learning algorithm based on Naive Bayes classification.
			</p>
			<button @click="fetch()" v-if="show_fetch">Fetch</button>
			<button @click="updateItems()" v-if="!show_fetch">Update Items</button>
			<div>
				<div v-for="item in items" class="article">
					<h1 @click="item.show_body = !item.show_body">{{item.data.title}}</h1>
					<p v-if="item.show_body">{{item.data.body}}</p>
				</div>
			</div>
		</div>
		<div class="page">
			<h1>STEP 2: Classify Articles</h1>
			<article >
				<h2>{{item.data.title}}</h2>
				<p>{{item.data.body}}</p>
			</article>
			<div>
				<div v-for="(score, cat) in item.data.categories">
					<input type="checkbox" :value="cat" :id="cat" v-model="item.categories">
					<label :for="cat">{{cat}}</label>
				</div>
				<input type="text" @keyup.enter="categories.push(newCategory)" v-model="newCategory" placeholder="Category"/>
				<button @click="submitClasses(); next()">Next</button>
			</div>
		</div>
	</div>
</template>
<script>
	export default {
	  data() {
	    return {
				query: "talimogene laherparepvec  [All Fields]",
	      threshold: 0.5,
	      newCategory: '',
				items: [],
				show_fetch: true,
				current_item_id: 0,
				show_score: 0,
	      item: {
	        data: {},
					id: 0,
	        categories: [],
	      },
				query: "talimogene laherparepvec [All Fields]"
	    }
	  },
	  created() {},
	  methods: {
			submitClasses() {
				this.$http.patch("/api/item/"+this.current_item_id+"/classes", {
					"classes": this.item.categories,
				})
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
						console.log(item)
						let item_wrapper = {
							show_body: false,
							data: item,
							categories: []
						}
						this.items.push(item_wrapper)
					}
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
				this.item = this.items[this.current_item_id]
				this.current_item_id++
	    },
			refreshCategory(category) {

			}
	  }
	}
</script>
