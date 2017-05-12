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
				<div v-for="item in all_items" class="article">
					<h1 @click="item.show_body = !item.show_body">{{item.data.title}}</h1>
					<p v-if="item.show_body">{{item.data.body}}</p>
				</div>
			</div>
		</div>
		<div class="page">
			<h1>STEP 2: Classify Articles</h1>
			<article >
				<h2>{{item.data.title}}</h2>
				<p>{{item.data.abstract}}</p>
			</article>
			<div>
				<div v-for="category in categories">
					<input type="checkbox" :value="category" :id="category" v-model="item.categories">
					<label :for="category">{{category}}</label>
				</div>
				<input type="text" @keyup.enter="categories.push(newCategory)" v-model="newCategory" placeholder="Category"/>
				<button @click="submitClasses(); next()">Next</button>
			</div>
		</div>
		<div class="page">
			<h1>STEP 3: Show Classified Articles</h1>
			<div v-for="category in categories">
				<h2>{{category}}</h2>
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
	      categories: ['relevant', 'irrelevent'],
				all_items: [],
				show_fetch: true,
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
				this.$http.patch("/api/item/"+this.item.id+"/classes", {
					"classes": this.item.categories,
				})
			},
			updateItems() {
				this.$http.get("/api/items").then(response => {
					this.$delete(this.all_items)
					let items = []
					for (let item of response.data.data) {
						this.all_items.push({
							show_body: false,
							data: item != "" ? item : "(no body)"
						})
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
	      this.$http.get("/act/item/next").then(response => {
					this.item.data = response.data.data;
					this.item.id = response.data.id;
	      }, response => {
					alert("There was an error getting the next item")
	        console(JSON.stringify(response));
	      });
	    },
			refreshCategory(category) {

			}
	  }
	}
</script>
