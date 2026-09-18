"""
Download the full GTEx v8 gene-level TPM matrix and the sample attributes
file, without filtering to any tissue -- so you can filter/subset yourself.

Files produced:
  - GTEx_Analysis_v8_gene_tpm.gct.gz   : full gene x sample TPM matrix (~1.6 GB)
  - GTEx_v8_SampleAttributes.txt       : per-sample metadata (tissue = SMTSD,
                                          plus sex, age bracket, RIN, etc. if
                                          you also pull the subject phenotypes file)

Requires: pip install pandas requests

Usage once downloaded:

    import gzip, pandas as pd

    attrs = pd.read_csv("GTEx_v8_SampleAttributes.txt", sep="\t")
    # e.g. attrs['SMTSD'].unique() to see all tissue labels

    # The .gct file has 2 metadata lines before the header row, so:
    df = pd.read_csv("GTEx_Analysis_v8_gene_tpm.gct.gz", sep="\t", skiprows=2)
    # df.columns[2:] are sample IDs matching attrs['SAMPID']

    my_samples = attrs.loc[attrs['SMTSD'] == 'Liver', 'SAMPID']
    liver_df = df[['Name', 'Description'] + [c for c in df.columns if c in my_samples.values]]
"""

import shutil
from pathlib import Path

import requests

GENE_TPM_URL = (
    "https://storage.googleapis.com/gtex_analysis_v8/rna_seq_data/"
    "GTEx_Analysis_2017-06-05_v8_RNASeQCv1.1.9_gene_tpm.gct.gz"
)
SAMPLE_ATTRS_URL = (
    "https://storage.googleapis.com/gtex_analysis_v8/annotations/"
    "GTEx_Analysis_v8_Annotations_SampleAttributesDS.txt"
)
# Optional: subject-level phenotypes (sex, age bracket, hardy scale) if you
# want to filter/stratify by donor characteristics too
SUBJECT_PHENOTYPES_URL = (
    "https://storage.googleapis.com/gtex_analysis_v8/annotations/"
    "GTEx_Analysis_v8_Annotations_SubjectPhenotypesDS.txt"
)

WORKDIR = Path(".")


def download(url: str, dest: Path):
    if dest.exists():
        print(f"Already downloaded: {dest} ({dest.stat().st_size / 1e6:.1f} MB)")
        return
    print(f"Downloading {url}\n  -> {dest}")
    with requests.get(url, stream=True, timeout=300) as r:
        r.raise_for_status()
        with open(dest, "wb") as f:
            shutil.copyfileobj(r.raw, f)
    print(f"  done ({dest.stat().st_size / 1e6:.1f} MB)")


if __name__ == "__main__":
    download(GENE_TPM_URL, WORKDIR / "GTEx_Analysis_v8_gene_tpm.gct.gz")
    download(SAMPLE_ATTRS_URL, WORKDIR / "GTEx_v8_SampleAttributes.txt")
    download(SUBJECT_PHENOTYPES_URL, WORKDIR / "GTEx_v8_SubjectPhenotypes.txt")

    print("\nAll files downloaded. Filter as needed -- see the module docstring")
    print("for a quick example of subsetting to a tissue with pandas.")












































